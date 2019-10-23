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
	ID          string
	StudentName string
	Urn         string
	ZipKey      string
	Status      Status
	CreatedTime *time.Time
}

type Status int

// Name represents a pseudo unique name that represent this task.
func (t *Task) Name() string {
	name := strings.ReplaceAll(t.StudentName, " ", "_")
	urn := t.Urn
	if len(urn) > 8 {
		urn = urn[:9]
	}

	return fmt.Sprintf("%s_%s", name, urn)
}
