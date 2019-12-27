package graderd

import (
	"context"
	"github.com/pkg/errors"
	"math/rand"
	"github.com/Capstone-auto-grader/grader-api-v2/internal/dockerdb"
	graderTask "github.com/Capstone-auto-grader/grader-api-v2/internal/grader-task"
)

type Scheduler interface {
	CreateImage(ctx context.Context, imageName string, imageTar []byte) error
	ListTasks(ctx context.Context) ([]*graderTask.Task, error)
	StartTask(ctx context.Context, taskID string) error
	EndTask(ctx context.Context, taskID string) error
	TaskOutput(ctx context.Context, taskID string) ([]byte, error)
}

type MockScheduler struct {
}

func NewMockScheduler() *MockScheduler {
	return &MockScheduler{}
}

func (m *MockScheduler) CreateImage(ctx context.Context, imageName string, imageTar []byte) error {
	return nil
}

func (m *MockScheduler) ListTasks(ctx context.Context, assignmentID string, db dockerdb.Database) ([]*graderTask.Task, error) {
	return nil, nil
}

func (m *MockScheduler) CreateTasks(ctx context.Context, taskList []*graderTask.Task, db dockerdb.Database) error {
	for _, t := range taskList {
		// generate "container id"
		containerID := make([]byte, 16)
		if _, err := rand.Read(containerID); err != nil {
			return errors.Wrap(err, ErrFailedToCreateTask.Error())
		}

		t.ContainerID = string(containerID)
		t.Status = graderTask.StatusPending
	}
	//err := db.PutTasks(ctx, taskList)
	//if err != nil {
	//	return err
	//}

	return nil
}

func (m *MockScheduler) StartTasks(ctx context.Context, taskList []*graderTask.Task, db dockerdb.Database) error {
	for _, task := range taskList {
		task.Status = graderTask.StatusStarted
		//if err := db.UpdateTask(ctx, task); err != nil {
		//	return err
		//}
	}

	return nil
}

func (m *MockScheduler) EndTask(ctx context.Context, taskID string, db dockerdb.Database) error {
	//task, err := db.GetTaskByID(ctx, taskID)
	//if err != nil {
	//	return err
	//}
	//task.Status = graderTask.StatusComplete
	//// update database
	//if err := db.UpdateTask(ctx, task); err != nil {
	//	return err
	//}

	return nil
}

func (m *MockScheduler) TaskOutput(ctx context.Context, taskID string, db dockerdb.Database) ([]byte, error) {
	sampleOutput := `
		this is a sample output generated by MockScheduler...
	`
	//if _, err := db.GetTaskByID(ctx, taskID); err != nil {
	//	return nil, ErrTaskNotFound
	//}

	return []byte(sampleOutput), nil
}
