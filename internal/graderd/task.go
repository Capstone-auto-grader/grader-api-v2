package graderd

import (
	"fmt"
	"strings"
	"time"
)

const (
	StatusPending Status = iota
	StatusStarted
	StatusComplete
	StatusFailed
)

// Task represents an assignment of a student.
type Task struct {
	// ID is a pseudo-unique name that represents this task.
	ID string
	// ContainerID represents the container that's running this task.
	ContainerID string
	// AssignmentID is the assignment that this task belongs to.
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
	t.ID = t.Name()
	return t
}

// Name is a pseudo unique name that represent this task.
func (t *Task) Name() string {
	name := strings.ReplaceAll(t.StudentName, " ", "_")

	return fmt.Sprintf("%s_%s_%s", name, t.AssignmentID, t.Urn)
}
