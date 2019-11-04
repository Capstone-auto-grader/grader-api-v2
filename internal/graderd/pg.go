package graderd

import (
	"context"
)

type PGClient struct {
}

func NewPGDatabase(databaseAddr string) Database {
	return &PGClient{}
}

func (PGClient) GetTaskByID(ctx context.Context, taskID string) (*Task, error) {
	panic("implement me")
}

func (PGClient) UpdateTask(ctx context.Context, task *Task) error {
	panic("implement me")
}

func (PGClient) PutTasks(ctx context.Context, taskList []*Task) error {
	panic("implement me")
}

func (PGClient) GetTasksByAssignment(ctx context.Context, assignmentID string) ([]*Task, error) {
	panic("implement me")
}
