package graderd

import (
	"context"
)

// Database is the interface that must be implemented by a database.
//
// Database implementation is encouraged to carry the given context.Context
// through out the whole execution process for better tracing.
type Database interface {
	// GetTaskByID retrieves task with an ID.
	//
	// Returns an error if task is not found.
	GetTaskByID(ctx context.Context, taskID string) (*Task, error)

	// GetTasksByAssignment retrieves tasks related/belongs to
	// the given assignment ID.
	//
	// Returns an error if the assignment is not found.
	GetTasksByAssignment(ctx context.Context, assignmentID string) ([]*Task, error)

	// UpdateTask updates the given task.
	UpdateTask(ctx context.Context, task *Task) error

	// PutTasks inserts/updates the given tasks in the database.
	PutTasks(ctx context.Context, taskList []*Task) error
}
