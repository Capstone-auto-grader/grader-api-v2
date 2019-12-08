package graderd

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/cenkalti/backoff/v3"
)

// ReturnResults return results back to $WEB_API.
// If an error is found returning result to host,
// it will retry with exponential backoff.
func (s *Service) ReturnResults() {
	for r := range s.results {
		b, err := json.Marshal(r)
		if err != nil {
			// We can silence the error because it is impossible to error here.
			b, _ = json.Marshal(struct {
				Error   error
				Message string
			}{
				Error:   err,
				Message: "Unable to return results, please retry.",
			})
		}
		// Retry with exponential backoff.
		_ = backoff.Retry(func() error {
			_, err = http.Post(s.webAddr, "application/json", bytes.NewReader(b))
			return err
		}, backoff.NewExponentialBackOff())
	}
}
