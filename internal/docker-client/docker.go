package docker_client

import (
	"bytes"
	"context"
	"fmt"
	"github.com/Capstone-auto-grader/grader-api-v2/internal/grader-task"
	"github.com/Capstone-auto-grader/grader-api-v2/internal/graderd"
	"github.com/Capstone-auto-grader/grader-api-v2/internal/sync-map"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"io/ioutil"
	"log"

	"github.com/pkg/errors"
)

type DockerClient struct {
	cli *client.Client
	mp *sync_map.SyncMap
	queue chan *grader_task.Task
}

// NewDockerClient creates a docker client for interacting with the docker host.
func NewDockerClient(host string, version string,  numWorkers int) *DockerClient {
	cli, err := client.NewClient(host, version, nil, nil)
	queue := make(chan *grader_task.Task)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	mp := sync_map.NewSyncMap()
	dockerClient := &DockerClient{
		cli: cli,
		mp: mp,
		queue: queue,
	}

	for i := 0; i < numWorkers; i++{
		go func() {
			WorkerEventLoop(queue, dockerClient, mp)
		}()
	}
	return dockerClient
}

// CreateAssignment with the given dockerfile and script, returns a unique assignment id.
func (d *DockerClient) CreateImage(ctx context.Context, imageName string, imageTar []byte) error {
	buildOptions := types.ImageBuildOptions{
		Dockerfile: "Dockerfile", // optional, is the default
		Tags:       []string{fmt.Sprintf("%s:latest", imageName)},
	}
	img := bytes.NewReader(imageTar)
	// build image
	_, err := d.cli.ImageBuild(ctx, img, buildOptions)
	if err != nil {
		return errors.Wrap(err, graderd.ErrFailedToBuildImage.Error())
	}

	return nil
}

// ListTasks lists all the active tasks associated with the assignment id in the docker host.
func (d *DockerClient) ListTasks(ctx context.Context) ([]*grader_task.Task, error) {
	containers, err := d.cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		return nil, err
	}

	tasks := d.mp.Enumerate()
	taskList := make([]*grader_task.Task, 0)

	for _, c := range containers {
		for _, t := range tasks {
			if t.ContainerID == c.ID {
				status := graderd.ParseContainerState(c.State)
				if status == grader_task.StatusStarted || status == grader_task.StatusFailed {
					taskList = append(taskList, &t)
				}
			}
		}
	}
	return taskList, nil
}

func (d *DockerClient) StartTask(ctx context.Context, task *grader_task.Task) error {
	d.queue <- task
	return nil
}
// EndTask stops the execution of the grader-task and remove its container from the host.
func (d *DockerClient) EndTask(ctx context.Context, task *grader_task.Task)error {
	return d.cli.ContainerRemove(ctx, task.ContainerID, types.ContainerRemoveOptions{})
}

func (d *DockerClient) StartContainerSync(ctx context.Context, task *grader_task.Task) error {
	return nil
}

func (d *DockerClient) createContainer(ctx context.Context, task *grader_task.Task) error {
	config := container.Config{
		AttachStdout:    true,
		AttachStderr:    true,
		Tty:             false,
		OpenStdin:       false,
		StdinOnce:       false,
		Env:             nil,
		Cmd:             nil,
		Healthcheck:     nil,
		ArgsEscaped:     false,
		Image:           "",
		Volumes:         nil,
		WorkingDir:      "",
		Entrypoint:      nil,
		NetworkDisabled: false,
		MacAddress:      "",
		OnBuild:         nil,
		Labels:          nil,
		StopSignal:      "",
		StopTimeout:     nil,
		Shell:           nil,
	}
}

// TaskOutput retrieves the stdout of the grader-task from the container.
func (d *DockerClient) TaskOutput(ctx context.Context, task *grader_task.Task) ([]byte, error) {
	task, err := d.mp.GetTask(task.ID)
	if err != nil {
		return nil, err
	}
	out, err := d.cli.ContainerLogs(ctx, task.ContainerID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(out)
}
