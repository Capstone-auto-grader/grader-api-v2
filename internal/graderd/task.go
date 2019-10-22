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

type Task struct {
	ID          string
	StudentName string
	Urn         string
	ZipKey      string
	Status      Status
	CreatedTime *time.Time
}

type Status int

func (t *Task) Name() string {
	name := strings.ReplaceAll(t.StudentName, " ", "_")
	urn := t.Urn
	if len(urn) > 8 {
		urn = urn[:9]
	}

	return fmt.Sprintf("%s_%s", name, urn)
}
