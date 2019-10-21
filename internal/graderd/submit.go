package graderd

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	graderpb "github.com/Capstone-auto-grader/grader-api-v2/graderpb"
)

// SubmitForGrading fulfills a SubmitForGradingRequest.
func (s *Service) SubmitForGrading(ctx context.Context, req *graderpb.SubmitForGradingRequest) (*graderpb.SubmitForGradingResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// TODO: Implement me
	return nil, status.Error(codes.Unimplemented, "not implemented")
}
