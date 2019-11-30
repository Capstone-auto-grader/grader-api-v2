package graderd

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

func (s *Service) ReturnResults(taskList []*Task) {
	wg := &sync.WaitGroup{}
	wg.Add(len(taskList))
	for _, task := range taskList {
		go func(task *Task) {
			t, err := s.schr.TaskOutput(context.Background(), task, s.db)
			if err != nil {
				log.Printf("failed to retreive output: %+v", err)
				return
			}
			b, err := json.Marshal(t)
			if err != nil {
				log.Printf("failed to marshal result: %+v", err)
				return
			}
			_, err = http.Post(s.webAddr, "application/json", bytes.NewReader(b))
			if err != nil {
				log.Printf("failed to return result: %+v", err)
				return
			}
			defer wg.Done()
		}(task)
	}
	wg.Wait()
}
