package dockerdb

import (
	"context"
	"database/sql"
	"github.com/Capstone-auto-grader/grader-api-v2/internal/grader-task"
)

// MemoryDB is an in-memory database.
type MemoryDB struct {
	assignmentIDs   []string
	assignmentTasks map[string][]*grader_task.Task
	tasksTable      map[string]*grader_task.Task
}

func NewMemoryDB() *MemoryDB {
	return &MemoryDB{
		assignmentIDs:   make([]string, 0),
		assignmentTasks: make(map[string][]*grader_task.Task),
		tasksTable:      make(map[string]*grader_task.Task),
	}
}

func (m *MemoryDB) GetTaskByID(ctx context.Context, taskID string) (*grader_task.Task, error) {
	t, ok := m.tasksTable[taskID]
	if !ok {
		return nil, sql.ErrNoRows
	}
	return t, nil
}

func (m *MemoryDB) UpdateTask(ctx context.Context, task *grader_task.Task) error {
	t, ok := m.tasksTable[task.Name()]
	if !ok {
		return sql.ErrNoRows
	}
	t.ID = task.ID
	t.Status = task.Status
	t.ContainerID = task.ContainerID
	t.CreatedTime = task.CreatedTime
	t.TestUri = task.TestUri
	t.SubmUri = task.SubmUri
	t.StudentName = task.StudentName
	t.AssignmentID = task.AssignmentID

	return nil
}

func (m *MemoryDB) PutTasks(ctx context.Context, taskList []*grader_task.Task) error {
	for _, t := range taskList {
		m.tasksTable[t.Name()] = t
		m.assignmentTasks[t.AssignmentID] = append(m.assignmentTasks[t.AssignmentID], t)
	}
	return nil
}

func (m *MemoryDB) GetTasksByAssignment(ctx context.Context, assignmentID string) ([]*grader_task.Task, error) {
	taskList, ok := m.assignmentTasks[assignmentID]
	if !ok {
		return nil, sql.ErrNoRows
	}

	return taskList, nil
}
