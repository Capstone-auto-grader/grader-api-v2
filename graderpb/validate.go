package grader

import "errors"

var (
	ErrCannotBeEmpty = errors.New("field cannot be empty")
)

// Validate validates a SubmitForGradingRequest.
// Currently, it only checks if all the fields are provided.
func (m *SubmitForGradingRequest) Validate() error {
	if m.GetTasks() == nil {
		return ErrCannotBeEmpty
	}

	for _, t := range m.GetTasks() {
		if err := t.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// Validate validates a Task.
// Currently, it only checks if all the fields are provided.
func (m *Task) Validate() error {
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

// Validate validates a CreateAssignmentRequest.
// Currently, it only checks if all the fields are provided.
func (m *CreateAssignmentRequest) Validate() error {
	if m.GetImageTar() == nil {
		return ErrCannotBeEmpty
	}

	return nil
}
