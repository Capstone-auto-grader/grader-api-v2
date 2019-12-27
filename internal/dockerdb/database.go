package dockerdb

import (
	"github.com/Capstone-auto-grader/grader-api-v2/internal/grader-task"
)


type Database interface {
	GetTaskByID(taskID string) (*grader_task.Task, error)
	GetTasksByAssignment(assignmentID string) ([]*grader_task.Task, error)
	UpdateTask(task *grader_task.Task) error
	PutTasks(taskList []*grader_task.Task) error
	UpdateStatus(taskID string, newStatus grader_task.Status) error
}
