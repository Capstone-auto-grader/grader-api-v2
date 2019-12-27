package docker_client

import (
	"bytes"
	"context"
	"fmt"
	"github.com/Capstone-auto-grader/grader-api-v2/internal/grader-task"
	"github.com/Capstone-auto-grader/grader-api-v2/internal/graderd"
	"github.com/Capstone-auto-grader/grader-api-v2/internal/sync-map"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"io/ioutil"
	"log"

	"github.com/pkg/errors"
)

type DockerClient struct {
	cli *client.Client
	mp *sync_map.SyncMap
}

// NewDockerClient creates a docker client for interacting with the docker host.
func NewDockerClient(host string, version string,mp *sync_map.SyncMap) *DockerClient {
	cli, err := client.NewClient(host, version, nil, nil)
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return &DockerClient{
		cli: cli,
		mp: mp,
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
		return errors.Wrap(err, graderd.ErrFailedToBuildImage.Error())
	}

	return nil
}

// ListTasks lists all the active tasks associated with the assignment id in the docker host.
func (d *DockerClient) ListTasks(ctx context.Context, assignmentID string) ([]grader_task.Task, error) {
	containers, err := d.cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		return nil, err
	}

	tasks := d.mp.Enumerate()
	taskList := make([]grader_task.Task, 0)

	for _, c := range containers {
		for _, t := range tasks {
			if t.ContainerID == c.ID {
				status := graderd.ParseContainerState(c.State)
				if status == grader_task.StatusStarted {
					taskList = append(taskList, t)
				}
			}
		}
	}
	return taskList, nil
}

//func (d *DockerClient) CreateTasks(ctx context.Context, taskList []*grader_task.Task, db dockerdb.Database) error {
//	// create tasks
//	for _, t := range taskList {
//		err := d.createTask(ctx, t)
//		if err != nil {
//			return err
//		}
//	}
//	// update database
//	if err := db.PutTasks(ctx, taskList); err != nil {
//		return errors.Wrap(err, ErrFailedToUpdateTask.Error())
//	}
//
//	return nil
//}
//
//func (d *DockerClient) createTask(ctx context.Context, task *grader_task.Task) error {
//	// create container
//	resp, err := d.cli.ContainerCreate(ctx, &container.Config{
//		Image: task.ImageID,
//		Labels: map[string]string{
//			"student_name": task.StudentName,
//		},
//		StopTimeout: task.Timeout,
//	}, nil, nil, task.ID)
//	if err != nil {
//		return errors.Wrap(err, ErrFailedToCreateTask.Error())
//	}
//
//	// assign id and time
//	task.ContainerID = resp.ID
//	t := time.Now()
//	task.CreatedTime = &t
//	task.Status = grader_task.StatusPending
//
//	return nil
//}

// StartTasks starts execution of all the given tasks (container IDs).
//func (d *DockerClient) StartTasks(ctx context.Context, taskList []*grader_task.Task, db dockerdb.Database) error {
//	for _, task := range taskList {
//		if err := d.cli.ContainerStart(ctx, task.ContainerID, types.ContainerStartOptions{}); err != nil {
//			return errors.Wrap(err, ErrFailedToStartTask.Error())
//		}
//		// update dockerdb
//		task.Status = grader_task.StatusStarted
//		if err := db.UpdateTask(ctx, task); err != nil {
//			return errors.Wrap(err, ErrFailedToUpdateTask.Error())
//		}
//	}
//	return nil
//}

func (d *DockerClient) StartTask(ctx context.Context, taskID string) error {
	return nil
}
// EndTask stops the execution of the grader-task and remove its container from the host.
func (d *DockerClient) EndTask(ctx context.Context, taskID string)error {
	task, err := d.mp.GetTask(taskID)
	if err != nil {
		return errors.Wrap(err, graderd.ErrTaskNotFound.Error())
	}

	return d.cli.ContainerRemove(ctx, task.ContainerID, types.ContainerRemoveOptions{})
}

// TaskOutput retrieves the stdout of the grader-task from the container.
func (d *DockerClient) TaskOutput(ctx context.Context, taskID string) ([]byte, error) {
	task, err := d.mp.GetTask(taskID)
	if err != nil {
		return nil, err
	}
	out, err := d.cli.ContainerLogs(ctx, task.ContainerID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(out)
}
