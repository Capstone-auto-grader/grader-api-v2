package graderd

import "context"

type Database interface {
	GetTaskByID(ctx context.Context, taskID string) (*Task, error)
	UpdateTask(ctx context.Context, task *Task) error
}
