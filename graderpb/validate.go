package grader

import "errors"

var (
	ErrCannotBeEmpty = errors.New("field cannot be empty")
)

// Validate validates a SubmitForGradingRequest.
// Currently, it only checks if all the fields are provided.
func (m *SubmitForGradingRequest) Validate() error {
	if len(m.GetUrnKey()) == 0 {
		return ErrCannotBeEmpty
	}

	if len(m.GetZipKey()) == 0 {
		return ErrCannotBeEmpty
	}

	if len(m.GetStudentName()) == 0 {
		return ErrCannotBeEmpty
	}

	return nil
}
