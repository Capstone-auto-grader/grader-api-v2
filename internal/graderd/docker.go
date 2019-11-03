package graderd

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"

	"github.com/pkg/errors"
)

type DockerClient struct {
	cli *client.Client
}

// NewDockerClient creates a docker client for interacting with the docker host.
func NewDockerClient(host string, version string) *DockerClient {
	cli, err := client.NewClient(host, version, nil, nil)
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return &DockerClient{
		cli: cli,
	}
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
		return errors.Wrap(err, ErrFailedToBuildImage.Error())
	}

	return nil
}

// ListTasks lists all the active tasks associated with the assignment id in the docker host.
func (d *DockerClient) ListTasks(ctx context.Context, assignmentID string, db Database) ([]*Task, error) {
	containers, err := d.cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		return nil, err
	}

	var taskList []*Task
	for _, c := range containers {
		t, err := db.GetTaskByID(ctx, c.ID)
		if err != nil {
			return nil, errors.Wrap(err, ErrTaskNotFound.Error())
		}
		t.Status = ParseContainerState(c.State)
		if err := db.UpdateTask(ctx, t); err != nil {
			return nil, err
		}
		taskList = append(taskList, t)
	}
	return taskList, nil
}

func (d *DockerClient) CreateTasks(ctx context.Context, taskList []*Task, db Database) ([]string, error) {
	// create tasks
	createdTasks := make([]string, 0, len(taskList))
	for _, t := range taskList {
		id, err := d.createTask(ctx, t)
		if err != nil {
			return nil, err
		}
		createdTasks = append(createdTasks, id)
	}
	// update database
	if err := db.PutTasks(ctx, taskList); err != nil {
		return nil, err
	}

	return createdTasks, nil
}

func (d *DockerClient) createTask(ctx context.Context, task *Task) (string, error) {
	// create container
	resp, err := d.cli.ContainerCreate(ctx, &container.Config{
		Image: task.AssignmentID,
		Labels: map[string]string{
			"student_name": task.StudentName,
		},
	}, nil, nil, task.ID)
	if err != nil {
		return "", errors.Wrap(err, ErrFailedToCreateTask.Error())
	}

	// assign id and time
	task.ContainerID = resp.ID
	t := time.Now()
	task.CreatedTime = &t

	return resp.ID, nil
}

// StartTasks starts execution of all the given tasks (container IDs).
func (d *DockerClient) StartTasks(ctx context.Context, taskIDs []string, db Database) error {
	for _, id := range taskIDs {
		if err := d.cli.ContainerStart(ctx, id, types.ContainerStartOptions{}); err != nil {
			return errors.Wrap(err, ErrFailedToStartTask.Error())
		}
		// update db
		t, err := db.GetTaskByID(ctx, id)
		if err != nil {
			return err
		}
		t.Status = StatusStarted
		if err := db.UpdateTask(ctx, t); err != nil {
			return err
		}
	}
	return nil
}

// EndTask stops the execution of the task and remove its container from the host.
func (d *DockerClient) EndTask(ctx context.Context, taskID string) error {
	return d.cli.ContainerRemove(ctx, taskID, types.ContainerRemoveOptions{})
}

// TaskOutput retrieves the stdout of the task from the container.
func (d *DockerClient) TaskOutput(ctx context.Context, id string) ([]byte, error) {
	out, err := d.cli.ContainerLogs(ctx, id, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(out)
}
