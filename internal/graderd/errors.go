package graderd

import (
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrAssignmentNotFound = errors.New("assignment not found")
	ErrTaskNotFound       = errors.New("grader-task not found")
	ErrFailedToCreateTask = errors.New("failed to create grader-task")
	ErrFailedToStartTask  = errors.New("failed to start grader-task")
	ErrFailedToUpdateTask = errors.New("failed to update grader-task")
	ErrFailedToBuildImage = errors.New("failed to build image")
)

func grpcError(err error) error {
	return status.Error(codes.Internal, err.Error())
}
