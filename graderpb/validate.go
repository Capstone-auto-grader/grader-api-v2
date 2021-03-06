package grader

import (
	"archive/tar"
	"bytes"
	"io"
	"regexp"

	"github.com/pkg/errors"
)

var (
	// errors
	ErrCannotBeEmpty     = errors.New("field cannot be empty")
	ErrInvalidFile       = errors.New("unable to parse file")
	ErrInvalidTimeout    = errors.New("invalid timeout: timeout has to be a positive non-zero number")
	ErrMissingDockerFile = errors.New("missing Dockerfile")
	ErrMissingRunScript  = errors.New("missing run script")
	// regex
	PatternDockerFile = regexp.MustCompile(`Dockerfile`)
	PatternRunScript  = regexp.MustCompile(`.\.sh`)
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

	if m.GetTimeout() <= 0 {
		return ErrInvalidTimeout
	}

	return nil
}

// Validate validates a CreateAssignmentRequest.
// Currently, it checks:
// - if it is a tar
// - if the tar contains a Dockerfile
// - if the tar contains a script
func (m *CreateAssignmentRequest) Validate() error {
	if m.GetImageName() == "" {
		return ErrCannotBeEmpty
	}

	if m.GetImageTar() == nil {
		return ErrCannotBeEmpty
	}

	tr := tar.NewReader(bytes.NewReader(m.GetImageTar()))
	dok, sok := false, false
	for {
		h, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return errors.Wrap(err, ErrInvalidFile.Error())
		}
		f := []byte(h.Name)
		if PatternDockerFile.Match(f) {
			dok = true
		}
		if PatternRunScript.Match(f) {
			sok = true
		}
	}
	// check if it has Dockerfile
	if !dok {
		return ErrMissingDockerFile
	}
	// check if it has run script
	if !sok {
		return ErrMissingRunScript
	}

	return nil
}
