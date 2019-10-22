package graderd

import "context"

type Scheduler interface {
	ListTasks(ctx context.Context) []*Task
	CreateTasks(ctx context.Context, image, imageURL string, taskList []*Task) ([]string, error)
	StartTasks(ctx context.Context, ids []string) error
	EndTask(ctx context.Context, id string) error

	TaskOutput(ctx context.Context, id string) ([]byte, error)
}
