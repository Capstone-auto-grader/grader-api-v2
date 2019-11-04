package graderd

import (
	"context"
	"database/sql"
)

// MemoryDB is an in-memory database.
type MemoryDB struct {
	assignmentIDs   []string
	assignmentTasks map[string][]*Task
	tasksTable      map[string]*Task
}

func NewMemoryDB() *MemoryDB {
	return &MemoryDB{
		assignmentIDs:   make([]string, 0),
		assignmentTasks: make(map[string][]*Task),
		tasksTable:      make(map[string]*Task),
	}
}

func (m *MemoryDB) GetTaskByID(ctx context.Context, taskID string) (*Task, error) {
	t, ok := m.tasksTable[taskID]
	if !ok {
		return nil, sql.ErrNoRows
	}
	return t, nil
}

func (m *MemoryDB) UpdateTask(ctx context.Context, task *Task) error {
	t, ok := m.tasksTable[task.Name()]
	if !ok {
		return sql.ErrNoRows
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

func (m *MemoryDB) PutTasks(ctx context.Context, taskList []*Task) error {
	for _, t := range taskList {
		m.tasksTable[t.Name()] = t
		m.assignmentTasks[t.AssignmentID] = append(m.assignmentTasks[t.AssignmentID], t)
	}
	return nil
}

func (m *MemoryDB) GetTasksByAssignment(ctx context.Context, assignmentID string) ([]*Task, error) {
	taskList, ok := m.assignmentTasks[assignmentID]
	if !ok {
		return nil, sql.ErrNoRows
	}

	return taskList, nil
}
