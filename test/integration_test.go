package test

import (
	"archive/tar"
	"bytes"
	"context"
	"fmt"
	"github.com/Capstone-auto-grader/grader-api-v2/internal/dockerdb"
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
		desc      string
		srv       *graderd.Service
		imageName string
		imageTar  []byte
		tasks     []*pb.Task
		err       error
	}{
		{
			desc:      "mock: one assignment, one submission",
			srv:       graderd.NewGraderService(graderd.NewMockScheduler(), dockerdb.NewMemoryDB(), ""),
			imageName: "assignment1",
			imageTar:  createHelloWorldImage(),
			tasks:     createNTasks(1),
			err:       nil,
		},
		{
			desc:      "mock: failed to create assignment (no image)",
			srv:       graderd.NewGraderService(graderd.NewMockScheduler(), dockerdb.NewMemoryDB(), ""),
			imageName: "assignment1",
			imageTar:  []byte{},
			tasks:     createNTasks(1),
			err:       status.Error(codes.InvalidArgument, pb.ErrMissingDockerFile.Error()),
		},
		{
			desc:      "mock: one assignment, one invalid submission",
			srv:       graderd.NewGraderService(graderd.NewMockScheduler(), dockerdb.NewMemoryDB(), ""),
			imageName: "assignment1",
			imageTar:  createValidImage(),
			tasks:     append(createNTasks(1), &pb.Task{}),
			err:       status.Error(codes.InvalidArgument, pb.ErrCannotBeEmpty.Error()),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			srv := test.srv
			ctx := context.Background()

			// create assignment
			creq := &pb.CreateAssignmentRequest{
				ImageName: test.imageName,
				ImageTar:  test.imageTar,
			}
			_, err := srv.CreateAssignment(ctx, creq)
			if err != nil {
				if want, got := test.err, err; !reflect.DeepEqual(want, got) {
					t.Errorf("unexpected error:\n- want: %+v\n- got: %+v\n", want, got)
				}
				return
			}
			// grade submissions
			greq := &pb.SubmitForGradingRequest{
				Tasks: fillAssignmentID(test.tasks, test.imageName),
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

func createValidImage() []byte {
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
			Timeout:      30,
		})
	}

	return tasks
}
