package graderd

import (
	"context"
)

type Database interface {
	GetTaskByID(ctx context.Context, taskID string) (*Task, error)
	UpdateTask(ctx context.Context, task *Task) error
	PutTasks(ctx context.Context, taskList []*Task) error
}

type MockDatabase struct {
	assignmentIDs   []string
	assignmentTasks map[string][]*Task
	tasksTable      map[string]*Task
}

func NewMockDatabase() *MockDatabase {
	return &MockDatabase{
		assignmentIDs:   make([]string, 0),
		assignmentTasks: make(map[string][]*Task),
		tasksTable:      make(map[string]*Task),
	}
}

func (m *MockDatabase) GetTaskByID(ctx context.Context, taskID string) (*Task, error) {
	t, ok := m.tasksTable[taskID]
	if !ok {
		return nil, ErrTaskNotFound
	}
	return t, nil
}

func (m *MockDatabase) UpdateTask(ctx context.Context, task *Task) error {
	t, ok := m.tasksTable[task.Name()]
	if !ok {
		return ErrTaskNotFound
	}
	t.ID = task.ID
	t.Status = task.Status
	t.ContainerID = task.ContainerID
	t.CreatedTime = task.CreatedTime
	t.ZipKey = task.ZipKey
	t.Urn = task.Urn
	t.StudentName = task.StudentName
	t.AssignmentID = task.AssignmentID

	return nil
}

func (m *MockDatabase) PutTasks(ctx context.Context, taskList []*Task) error {
	for _, t := range taskList {
		m.tasksTable[t.Name()] = t
	}
	return nil
}
