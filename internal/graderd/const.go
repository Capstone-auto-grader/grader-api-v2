package graderd

import "github.com/pkg/errors"

var (
	ErrAssignmentNotFound = errors.New("assignment not found")
	ErrTaskNotFound       = errors.New("task not found")
)
