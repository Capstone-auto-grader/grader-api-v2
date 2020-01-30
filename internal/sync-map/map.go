package sync_map

import (
	"fmt"
	"sync"
	"github.com/Capstone-auto-grader/grader-api-v2/internal/grader-task"
)
// A synchronized map to store Tasks
//
// Note-- this is the canonical location for
// Tasks to be stored. Given that they are
// passed as values, not as reference
// TODO: Error handling for not found
type SyncMap struct {
	mu *sync.RWMutex
	mp map[string]*grader_task.Task
}

func NewSyncMap() *SyncMap{
	return &SyncMap{
		mu: &sync.RWMutex{},
		mp: make(map[string]*grader_task.Task),
	}
}

// TODO: decide where to handle duplicate requests
func (m *SyncMap) StoreTask(task *grader_task.Task) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.mp[task.ID] = task
}

func (m *SyncMap) GetTask(taskID string) (*grader_task.Task, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if val, ok := m.mp[taskID]; ok {
		return val, nil
	} else {
		return nil, fmt.Errorf("Not found")
	}
}

// TODO: make sure that duplicate requests don't prompt anomalous state
// TODO: Just deal with pointers
func (m *SyncMap) UpdateTaskContainerID(taskID string, containerId string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	t := m.mp[taskID]
	t.ContainerID = containerId
	m.mp[taskID] = t
}

//TODO: JUST DEAL WITH POINTERS
func (m *SyncMap) UpdateStatus(taskID string, status grader_task.Status, checkStatus bool) error{
	m.mu.Lock()
	defer m.mu.Unlock()
	t, ok := m.mp[taskID]
	if ok {
		if checkStatus && t.Status == grader_task.StatusStarted {
			return fmt.Errorf("task already started")
		}
		t.Status = status
	}

	m.mp[taskID] = t
	return nil
}

func (m *SyncMap) Enumerate() map[string]grader_task.Task {
	m.mu.RLock()
	defer m.mu.RUnlock()

	// Given that we are dealing with multiple workers,
	// it is safest to simply make a copy of the map
	// when we want to do something like check the status.
	// It is not an issue to have slightly out-of-date information
	// as long as the integrity of the internal map is preserved
	ret := make(map[string]grader_task.Task)
	for k,v := range m.mp {
		newval := *v
		ret[k] = newval
	}
	return ret
}
