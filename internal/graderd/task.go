package graderd

import (
	"strings"
	"time"
)

const (
	StatusPending Status = iota
	StatusStarted
	StatusComplete
	StatusFailed
)

// Task represents a student's assignment.
type Task struct {
	// ID is a pseudo-unique name that represents this task.
	ID string

	// ContainerID represents the container that's running this task.
	ContainerID string

	// AssignmentID represents the assignment that this task belongs to.
	AssignmentID string
	StudentName  string
	Urn          string
	ZipKey       string

	// Timeout to stop a container in seconds.
	Timeout *int
	Status  Status

	// CreatedTime is the time when the container is created for this task.
	CreatedTime *time.Time
}

// Status represents a task's status.
type Status int

// NewTask creates a task.
func NewTask(assignmentID, urn, zip, studentName string, timeout int32) *Task {
	to := int(timeout)
	t := &Task{
		AssignmentID: assignmentID,
		Urn:          urn,
		ZipKey:       zip,
		StudentName:  studentName,
		Timeout:      &to,
	}
	t.ID = t.id()
	return t
}

// id generates a pseudo unique name representing this task.
func (t *Task) id() string {
	name := strings.ReplaceAll(t.StudentName, " ", "_")

	return strings.Join([]string{name, t.AssignmentID, t.Urn}, "_")
}
