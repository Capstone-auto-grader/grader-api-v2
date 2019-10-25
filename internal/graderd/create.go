package graderd

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/Capstone-auto-grader/grader-api-v2/graderpb"
)

func (s *Service) CreateAssignment(ctx context.Context, req *pb.CreateAssignmentRequest) (*pb.CreateAssignmentResponse, error) {
	// TODO: Implement me
	return nil, status.Error(codes.Unimplemented, "not implemented")
}
