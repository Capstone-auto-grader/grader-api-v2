package graderd

import (
	"context"
	"github.com/Capstone-auto-grader/grader-api-v2/internal/grader-task"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/Capstone-auto-grader/grader-api-v2/graderpb"
)

// SubmitForGrading fulfills a SubmitForGradingRequest.
func (s *Service) SubmitForGrading(ctx context.Context, req *pb.SubmitForGradingRequest) (*pb.SubmitForGradingResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// Create tasks.

	for _, t := range req.GetTasks() {
		task := grader_task.NewTask(t.GetImageName(), t.GetZipKey(), t.GetTestKey(), t.GetStudentName(), t.GetCallbackUri(), t.GetTimeout())
		_ = s.schd.StartTask(ctx, &task)
	}


	//err := s.schr.CreateTasks(ctx, taskList, s.db)
	//if err != nil {
	//	return nil, grpcError(err)
	//}
	//// Start tasks.
	//if err := s.schr.StartTasks(ctx, taskList, s.db); err != nil {
	//	return nil, grpcError(err)
	//}

	return &pb.SubmitForGradingResponse{}, nil
}

func (s *Service) CreateImage(ctx context.Context, req *pb.CreateImageRequest) (*pb.CreateImageResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err := s.schd.CreateImage(ctx, req.GetImageName(), req.GetImageTar())
	if err != nil {
		return nil, grpcError(err)
	}

	return &pb.CreateImageResponse{}, nil
}
