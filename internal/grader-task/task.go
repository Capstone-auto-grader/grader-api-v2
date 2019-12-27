package grader_task

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
// Tasks should ONLY BE PASSED BY VALUE in order to
// avoid synchronization problems
type Task struct {
	// ID is a pseudo-unique name that represents this grader-task.
	ID string
	// ContainerID represents the container that's running this grader-task.
	ContainerID string
	// AssignmentID is the assignment that this grader-task belongs to.
	ImageID		 string
	StudentName  string
	SubmUri      string
	TestUri      string
	CallbackUri  string
	// Timeout to stop a container in seconds.
	Timeout *int
	Status  Status
	// CreatedTime is the time when the container is created for this grader-task.
	CreatedTime *time.Time
}

type Status int

// NewTask creates a grader-task.
func NewTask(imageID, submUri, testZip, studentName, callbackUri string, timeout int32) Task {
	to := int(timeout)
	t := Task{
		ImageID: 	  imageID,
		SubmUri:      submUri,
		TestUri:      testZip,
		CallbackUri: callbackUri,
		StudentName:  studentName,
		Timeout:      &to,
	}
	t.ID = t.Name()
	return t
}

// Name is a pseudo unique name that represent this grader-task.
func (t *Task) Name() string {
	name := strings.ReplaceAll(t.StudentName, " ", "_")

	return fmt.Sprintf("%s_%s_%s", name, t.ImageID, t.SubmUri)
}
