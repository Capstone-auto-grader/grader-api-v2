package graderd

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/cenkalti/backoff/v3"
)

// ReturnResults return results back to $WEB_API.
// If an error is found returning result to host,
// it will retry with exponential backoff.
func (s *Service) ReturnResults(taskList []*Task) {
	wg := &sync.WaitGroup{}

	wg.Add(len(taskList))
	for _, task := range taskList {
		go func(task *Task) {
			ctx := context.Background()

			// Retrieve task's output.
			t, err := s.schr.TaskOutput(ctx, task, s.db)
			if err != nil {
				log.Printf("failed to retreive output: %+v", err)
				return
			}

			// Marshal Output.
			b, err := json.Marshal(t)
			if err != nil {
				log.Printf("failed to marshal result: %+v", err)
				return
			}

			// Retry with exponential backoff.
			_ = backoff.Retry(func() error {
				_, err = http.Post(s.webAddr, "application/json", bytes.NewReader(b))
				return err
			}, backoff.NewExponentialBackOff())

			wg.Done()
		}(task)
	}
	// Wait for all tasks to finish.
	wg.Wait()
}
