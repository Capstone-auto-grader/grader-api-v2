package graderd

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (s *Service) ReturnResults() {
	for r := range s.results {
		b, err := json.Marshal(r)
		if err != nil {

		}
		_, err = http.Post(s.webAddr, "application/json", bytes.NewReader(b))
		if err != nil {

		}
	}
}
