package worker

import (
	"context"
	"github.com/Capstone-auto-grader/grader-api-v2/internal/docker-client"
	"github.com/Capstone-auto-grader/grader-api-v2/internal/grader-task"
	sync_map "github.com/Capstone-auto-grader/grader-api-v2/internal/sync-map"
)

func WorkerEventLoop(queue <-chan string, client *docker_client.DockerClient, tasks *sync_map.SyncMap) {
	for {
		taskId := <-queue
		_ = tasks.UpdateStatus(taskId, grader_task.StatusStarted, true)
		_ = client.StartTask(context.Background(), taskId)
		_,_ = client.TaskOutput(context.Background(), taskId)
	}
}