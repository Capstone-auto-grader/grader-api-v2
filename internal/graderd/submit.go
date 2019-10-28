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

	// Create tasks.
	taskList := make([]*Task, 0, len(req.GetTasks()))
	for _, t := range req.GetTasks() {
		task := NewTask(t.GetAssignmentId(), t.GetUrnKey(), t.GetZipKey(), t.GetStudentName())
		taskList = append(taskList, task)
	}
	// TODO: Add image and image url
	ids, err := s.schr.CreateTasks(ctx, "", "", taskList)
	if err != nil {
		return nil, grpcError(err)
	}
	// Start tasks.
	if err := s.schr.StartTasks(ctx, ids, nil); err != nil {
		return nil, grpcError(err)
	}

	return &graderpb.SubmitForGradingResponse{}, nil
}
