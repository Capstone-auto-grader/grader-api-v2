package graderd

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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
	// Output is the raw output of the container job.
	Output io.ReadCloser
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

func (t *Task) MarshalJSON() ([]byte, error) {
	b, err := ioutil.ReadAll(t.Output)
	if err != nil {
		return nil, err
	}
	defer t.Output.Close()

	return json.Marshal(&struct {
		ID           string
		AssignmentID string
		Output       []byte
	}{
		ID:           t.ID,
		AssignmentID: t.AssignmentID,
		Output:       b,
	})
}
