package graderd

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/Capstone-auto-grader/grader-api-v2/graderpb"
)

// CreateAssignment creates an assignment using the given parameters through the scheduler.
func (s *Service) CreateAssignment(ctx context.Context, req *pb.CreateAssignmentRequest) (*pb.CreateAssignmentResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err := s.schr.CreateImage(ctx, req.GetImageName(), req.GetImageTar())
	if err != nil {
		return nil, gRPCError(err)
	}

	return &pb.CreateAssignmentResponse{}, nil
}
