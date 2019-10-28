package graderd

import (
	"context"
	"crypto/md5"
)

type Scheduler interface {
	CreateAssignment(ctx context.Context, dockerFile []byte, script []byte) (string, error)

	ListTasks(ctx context.Context, assignmentID string) []*Task
	CreateTasks(ctx context.Context, image, imageURL string, taskList []*Task) ([]string, error)
	StartTasks(ctx context.Context, taskIDs []string) error
	EndTask(ctx context.Context, taskID string) error

	TaskOutput(ctx context.Context, id string) ([]byte, error)
}

type MockScheduler struct {
	assignmentIDs   []string
	assignmentTasks map[string][]*Task
	tasksTable      map[string]*Task
}

func (m *MockScheduler) CreateAssignment(ctx context.Context, dockerFile []byte, script []byte) (string, error) {
	h := md5.Sum(append(dockerFile, script...))
	id := string(h[:])
	m.assignmentIDs = append(m.assignmentIDs, id)
	return id, nil
}

func (m *MockScheduler) ListTasks(ctx context.Context, assignmentID string) []*Task {
	return m.assignmentTasks[assignmentID]
}

func (m *MockScheduler) existsAssignment(ctx context.Context, assignmentID string) bool {
	for _, id := range m.assignmentIDs {
		if id == assignmentID {
			return true
		}
	}
	return false
}

func (m *MockScheduler) CreateTasks(ctx context.Context, image, imageURL string, taskList []*Task) ([]string, error) {
	ids := make([]string, 0, len(taskList))
	for _, t := range taskList {
		if !m.existsAssignment(ctx, t.AssignmentID) {
			return nil, ErrAssignmentNotFound
		}

		// generate "container id"
		id := md5.Sum([]byte(t.Name()))
		t.ID = string(id[:])
		ids = append(ids, string(id[:]))

		// store tasks in assignments
		m.assignmentTasks[t.AssignmentID] = append(m.assignmentTasks[t.AssignmentID], t)
		m.tasksTable[t.ID] = t
	}

	return ids, nil
}

func (m *MockScheduler) StartTasks(ctx context.Context, taskIDs []string) error {
	for _, id := range taskIDs {
		t, ok := m.tasksTable[id]
		if !ok {
			return ErrTaskNotFound
		}
		t.Status = StatusStarted
	}

	return nil
}

func (m *MockScheduler) EndTask(ctx context.Context, taskID string) error {
	if t, ok := m.tasksTable[taskID]; !ok {
		return ErrTaskNotFound
	} else {
		t.Status = StatusComplete
	}
	return nil
}

func (m *MockScheduler) TaskOutput(ctx context.Context, id string) ([]byte, error) {
	sampleOutput := `
		this is a sample output...
	`
	if _, ok := m.tasksTable[id]; !ok {
		return nil, ErrTaskNotFound
	}

	return []byte(sampleOutput), nil
}
