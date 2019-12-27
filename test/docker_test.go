package test

import (
	"context"
	"github.com/Capstone-auto-grader/grader-api-v2/internal/docker-client"
	db2 "github.com/Capstone-auto-grader/grader-api-v2/internal/dockerdb"
	"github.com/Capstone-auto-grader/grader-api-v2/internal/grader-task"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestDockerClient_CreateAssignment(t *testing.T) {
	tests := []struct {
		desc      string
		host      string
		version   string
		imageName string
		imageTar  []byte
		err       error
	}{
		{
			desc:      "create hello world",
			host:      "http://localhost:2376",
			version:   "1.40",
			imageName: "helloworld",
			imageTar:  createHelloWorldImage(),
			err:       nil,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			d := docker_client.NewDockerClient(test.host, test.version)
			ctx := context.Background()
			// create assignment
			err := d.CreateImage(ctx, test.imageName, test.imageTar)
			if err != nil {
				if want, got := test.err, err; !reflect.DeepEqual(want, got) {
					t.Errorf("unexpected error:\n- want: %+v\n- got: %+v\n", want, got)
				}
				return
			}
		})
	}
}

func TestDockerClient_CreateTasks(t *testing.T) {
	tests := []struct {
		desc      string
		host      string
		version   string
		imageName string
		tasks     []*grader_task.Task
		err       error
	}{
		{
			desc:      "one grader-task",
			host:      "http://localhost:2376",
			version:   "1.40",
			imageName: "helloworld",
			tasks: []*grader_task.Task{
				{
					AssignmentID: "helloworld",
					StudentName:  "some name",
					SubmUri:      "some:urn:key",
					TestUri:      "some:zip:key",
				},
			},
			err: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			d := docker_client.NewDockerClient(test.host, test.version)
			db := db2.NewMemoryDB()
			ctx := context.Background()
			// create tasks
			err := d.CreateTasks(ctx, test.tasks, db)
			if err != nil {
				if want, got := test.err, err; !reflect.DeepEqual(want, got) {
					t.Errorf("unexpected error:\n- want: %+v\n- got: %+v\n", want, got)
				}
				return
			}
		})
	}
}

func TestDockerClient_CreateAndStartTasks(t *testing.T) {
	tests := []struct {
		desc      string
		host      string
		version   string
		imageName string
		tasks     []*grader_task.Task
		err       error
	}{
		{
			desc:      "one grader-task, foreign image",
			host:      "http://localhost:2376",
			version:   "1.40",
			imageName: "hello-world",
			tasks: []*grader_task.Task{
				{
					AssignmentID: "hello-world",
					StudentName:  "some name",
					SubmUri:      "some:urn:key",
					TestUri:      "some:zip:key",
				},
			},
			err: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			d := docker_client.NewDockerClient(test.host, test.version)
			db := db2.NewMemoryDB()
			ctx := context.Background()
			// create tasks
			err := d.CreateTasks(ctx, test.tasks, db)
			if err != nil {
				if want, got := test.err, err; !reflect.DeepEqual(want, got) {
					t.Errorf("unexpected error:\n- want: %+v\n- got: %+v\n", want, got)
				}
				return
			}
			// start tasks
			err = d.StartTasks(ctx, test.tasks, db)
			if err != nil {
				if want, got := test.err, err; !reflect.DeepEqual(want, got) {
					t.Errorf("unexpected error:\n- want: %+v\n- got: %+v\n", want, got)
				}
				return
			}
		})
	}
}

func createHelloWorldImage() []byte {
	b, err := ioutil.ReadFile("helloworld.tar.gz")
	if err != nil {
		panic(err)
	}
	return b
}
