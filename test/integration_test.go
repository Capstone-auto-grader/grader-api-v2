package test

import (
	"archive/tar"
	"bytes"
	"context"
	"fmt"
	"log"
	"reflect"
	"testing"

	pb "github.com/Capstone-auto-grader/grader-api-v2/graderpb"
	"github.com/Capstone-auto-grader/grader-api-v2/internal/graderd"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCreateAssignmentAndGrade(t *testing.T) {
	tests := []struct {
		desc     string
		imageTar []byte
		tasks    []*pb.Task
		err      error
	}{
		{
			desc:     "one assignment, one submission",
			imageTar: createImage(),
			tasks:    createNTasks(1),
			err:      nil,
		},
		{
			desc:     "failed to create assignment (no image)",
			imageTar: []byte{},
			tasks:    createNTasks(1),
			err:      status.Error(codes.InvalidArgument, pb.ErrMissingDockerFile.Error()),
		},
	}

	for _, test := range tests {
		srv := graderd.NewGraderService(graderd.NewMockScheduler(), "")
		ctx := context.Background()

		t.Run(test.desc, func(t *testing.T) {
			// create assignment
			creq := &pb.CreateAssignmentRequest{
				ImageTar: test.imageTar,
			}
			cres, err := srv.CreateAssignment(ctx, creq)
			if err != nil {
				if want, got := test.err, err; !reflect.DeepEqual(want, got) {
					t.Errorf("unexpected error:\n- want: %+v\n- got: %+v\n", want, got)
				}
				return
			}
			id := cres.GetId()
			if id == "" {
				t.Errorf("unexpected empty assignmentID: %+v\n", id)
			}
			// grade submissions
			greq := &pb.SubmitForGradingRequest{
				Tasks: fillAssignmentID(test.tasks, id),
			}
			_, err = srv.SubmitForGrading(ctx, greq)
			if err != nil {
				if want, got := test.err, err; !reflect.DeepEqual(want, got) {
					t.Errorf("unexpected error:\n- want: %+v\n- got: %+v\n", want, got)
				}
			}
		})
	}
}

func fillAssignmentID(tasks []*pb.Task, id string) []*pb.Task {
	for _, t := range tasks {
		t.AssignmentId = id
	}
	return tasks
}

func createImage() []byte {
	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)
	var files = []struct {
		Name, Body string
	}{
		{"Dockerfile", "This is a Dockerfile."},
		{"run.sh", "This is a run script."},
	}
	for _, file := range files {
		hdr := &tar.Header{
			Name: file.Name,
			Mode: 0600,
			Size: int64(len(file.Body)),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			log.Fatal(err)
		}
		if _, err := tw.Write([]byte(file.Body)); err != nil {
			log.Fatal(err)
		}
	}

	return buf.Bytes()
}

func createNTasks(n int) []*pb.Task {
	tasks := make([]*pb.Task, 0, n)

	for i := 0; i < n; i++ {
		tasks = append(tasks, &pb.Task{
			AssignmentId: "",
			UrnKey:       fmt.Sprintf("some:%d:urn:key", i),
			ZipKey:       fmt.Sprintf("some:%d:zip:key", i),
			StudentName:  fmt.Sprintf("some kid %d", i),
		})
	}

	return tasks
}
