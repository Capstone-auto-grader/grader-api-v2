package graderd

import (
	"context"
)

type Database interface {
	GetTaskByID(ctx context.Context, taskID string) (*Task, error)
	GetTasksByAssignment(ctx context.Context, assignmentID string) ([]*Task, error)
	UpdateTask(ctx context.Context, task *Task) error
	PutTasks(ctx context.Context, taskList []*Task) error
}
