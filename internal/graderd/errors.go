package graderd

import (
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrAssignmentNotFound = errors.New("assignment not found")
	ErrTaskNotFound       = errors.New("task not found")
	ErrFailedToBuildImage = errors.New("failed to build image")
)

func grpcError(err error) error {
	return status.Error(codes.Internal, err.Error())
}
