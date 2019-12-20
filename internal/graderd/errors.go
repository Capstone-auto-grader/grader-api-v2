package graderd

import (
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	// ErrAssignmentNotFound represents an error finding the assignment.
	ErrAssignmentNotFound = errors.New("assignment not found")

	// ErrTaskNotFound represents an error finding the task.
	ErrTaskNotFound = errors.New("task not found")

	// ErrFailedToCreateTask represents an error creating a task.
	ErrFailedToCreateTask = errors.New("failed to create task")

	// ErrFailedToStartTask represents an error starting a given task.
	ErrFailedToStartTask = errors.New("failed to start task")

	// ErrFailedToUpdateTask represents an error updating a given task.
	ErrFailedToUpdateTask = errors.New("failed to update task")

	// ErrFailedToBuildImage represents an error building an image.
	ErrFailedToBuildImage = errors.New("failed to build image")
)

// gRPCError wraps an error with a gRPC status code.
//
// Currently, it wraps every error with INTERNAL(13) code that signals
// a non-retryable internal error.
// See: https://github.com/grpc/grpc/blob/master/doc/statuscodes.md
//
// In the future, we will provide more fine grained code status by unwrapping
// and identifying the underlying error.
func gRPCError(err error) error {
	return status.Error(codes.Internal, err.Error())
}
