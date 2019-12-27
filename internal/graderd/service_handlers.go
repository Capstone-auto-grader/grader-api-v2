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
	taskList := make([]grader_task.Task, 0, len(req.GetTasks()))
	for _, t := range req.GetTasks() {
		task := grader_task.NewTask(t.GetAssignmentId(), t.GetZipKey(), t.GetTestKey(), t.GetStudentName(), t.GetCallbackUri(), t.GetTimeout())
		taskList = append(taskList, task)
		s.mp.StoreTask(&task)
		s.jobChan <- task.ID
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

func (s *Service) CreateAssignment(ctx context.Context, req *pb.CreateAssignmentRequest) (*pb.CreateAssignmentResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err := s.schd.CreateImage(ctx, req.GetImageName(), req.GetImageTar())
	if err != nil {
		return nil, grpcError(err)
	}

	return &pb.CreateAssignmentResponse{}, nil
}
