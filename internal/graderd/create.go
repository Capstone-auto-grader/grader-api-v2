package graderd

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/Capstone-auto-grader/grader-api-v2/graderpb"
)

// CreateAssignment creates an assignment using a given dockerfile and run script.
func (s *Service) CreateAssignment(ctx context.Context, req *pb.CreateAssignmentRequest) (*pb.CreateAssignmentResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	id, err := s.schr.CreateAssignment(ctx, req.GetImageTar())
	if err != nil {
		return nil, grpcError(err)
	}

	return &pb.CreateAssignmentResponse{
		Id: id,
	}, nil
}
