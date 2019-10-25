package graderd

import "context"

type Scheduler interface {
	CreateAssignment(ctx context.Context, dockerFile []byte, script []byte) (string, error)

	ListTasks(ctx context.Context, assignmentID string) []*Task
	CreateTasks(ctx context.Context, image, imageURL string, taskList []*Task) ([]string, error)
	StartTasks(ctx context.Context, ids []string) error
	EndTask(ctx context.Context, id string) error

	TaskOutput(ctx context.Context, id string) ([]byte, error)
}
