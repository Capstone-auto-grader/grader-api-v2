package graderd

import (
	"fmt"
	"strings"
	"time"
)

const (
	Pending Status = iota
	Complete
	Failed
)

// Task represents an assignment of a student.
type Task struct {
	ID           string
	AssignmentID string
	StudentName  string
	Urn          string
	ZipKey       string
	Status       Status
	CreatedTime  *time.Time
}

type Status int

// Name is a pseudo unique name that represent this task.
func (t *Task) Name() string {
	name := strings.ReplaceAll(t.StudentName, " ", "_")

	return fmt.Sprintf("%s_%s_%s", name, t.AssignmentID, t.Urn)
}
