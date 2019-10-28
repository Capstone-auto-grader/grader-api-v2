package graderd

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type DockerClient struct {
	cli *client.Client
}

// NewDockerClient creates a docker client for interacting with the docker host.
func NewDockerClient(host string, version string) *DockerClient {
	cli, err := client.NewClient(host, version, &http.Client{}, nil)
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return &DockerClient{
		cli: cli,
	}
}

// CreateAssignment with the given dockerfile and script, returns a unique assignment id.
func (d *DockerClient) CreateAssignment(ctx context.Context, dockerFile []byte, script []byte) (string, error) {

	return "", nil
}

// ListTasks lists all the active tasks associated with the assignment id in the docker host.
func (d *DockerClient) ListTasks(ctx context.Context, assignmentID string) []*Task {

	return nil
}

func (d *DockerClient) CreateTasks(ctx context.Context, image, imageURL string, taskList []*Task) ([]string, error) {
	// pull image
	_, err := d.cli.ImagePull(ctx, imageURL, types.ImagePullOptions{})
	if err != nil {
		return nil, err
	}

	// create tasks
	createdTasks := make([]string, 0, len(taskList))
	for _, t := range taskList {
		id, err := d.createTask(ctx, image, t)
		if err != nil {
			return nil, err
		}
		createdTasks = append(createdTasks, id)
	}

	return createdTasks, nil
}

func (d *DockerClient) createTask(ctx context.Context, image string, task *Task) (string, error) {
	// create container
	resp, err := d.cli.ContainerCreate(ctx, &container.Config{
		Image: image,
		Labels: map[string]string{
			"assignment_id": task.AssignmentID,
			"student_name":  task.StudentName,
		},
	}, nil, nil, task.Name())
	if err != nil {
		return "", err
	}

	// assign id
	task.ID = resp.ID
	return resp.ID, nil
}

// StartTasks starts execution of all the given tasks (container IDs).
func (d *DockerClient) StartTasks(ctx context.Context, taskIDs []string) error {
	for _, id := range taskIDs {
		if err := d.cli.ContainerStart(ctx, id, types.ContainerStartOptions{}); err != nil {
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
