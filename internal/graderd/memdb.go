package graderd

import (
	"context"

	"github.com/hashicorp/go-memdb"
)

type MemDB struct {
	*memdb.MemDB
}

func NewMemDB() *MemDB {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"task": {
				Name: "task",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"containerID": {
						Name:    "containerID",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "ContainerID"},
					},
					"assignmentID": {
						Name:    "assignmentID",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "AssignmentID"},
					},
				},
			},
		},
	}
	db, err := memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}

	return &MemDB{
		db,
	}
}

func (m *MemDB) GetTaskByID(ctx context.Context, taskID string) (*Task, error) {
	tx := m.Txn(false)
	defer tx.Abort()

	row, err := tx.First("task", "id", taskID)
	if err != nil {
		return nil, err
	}
	return row.(*Task), nil
}

func (m *MemDB) GetTasksByAssignment(ctx context.Context, assignmentID string) ([]*Task, error) {
	tx := m.Txn(false)
	defer tx.Abort()

	var tasks []*Task
	rows, err := tx.Get("task", "assignmentID")
	if err != nil {
		return nil, err
	}
	for row := rows.Next(); row != nil; row = rows.Next() {
		tasks = append(tasks, row.(*Task))
	}
	return tasks, nil
}

func (m *MemDB) UpdateTask(ctx context.Context, task *Task) error {
	tx := m.Txn(true)
	defer tx.Commit()

	return tx.Insert("task", task)
}

func (m *MemDB) PutTasks(ctx context.Context, taskList []*Task) error {
	tx := m.Txn(true)
	defer tx.Commit()

	for _, task := range taskList {
		if err := tx.Insert("task", task); err != nil {
			return err
		}
	}
	return nil
}
