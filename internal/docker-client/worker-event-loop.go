package docker_client

import (
	"context"
	"github.com/Capstone-auto-grader/grader-api-v2/internal/grader-task"
	sync_map "github.com/Capstone-auto-grader/grader-api-v2/internal/sync-map"
	"log"
)

func WorkerEventLoop(queue <-chan *grader_task.Task, client *DockerClient, tasks *sync_map.SyncMap) {
	for {
		task:= <-queue
		if err := tasks.UpdateStatus(task.ID, grader_task.StatusStarted, true); err != nil {
			log.Println(err.Error())
		}
		if err := client.StartContainerSync(context.Background(), task); err != nil {
			log.Println(err.Error())
		}
		out,err := client.TaskOutput(context.Background(), task)
		if err != nil {
			log.Println(err.Error())
		}
		log.Println(string(out))
	}
}