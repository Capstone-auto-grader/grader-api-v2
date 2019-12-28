package docker_client

import (
	"context"
	"github.com/Capstone-auto-grader/grader-api-v2/internal/grader-task"
	sync_map "github.com/Capstone-auto-grader/grader-api-v2/internal/sync-map"
)

func WorkerEventLoop(queue <-chan *grader_task.Task, client *DockerClient, tasks *sync_map.SyncMap) {
	for {
		task:= <-queue
		_ = tasks.UpdateStatus(task.ID, grader_task.StatusStarted, true)
		_ = client.StartContainerSync(context.Background(), task)
		_,_ = client.TaskOutput(context.Background(), task)
	}
}