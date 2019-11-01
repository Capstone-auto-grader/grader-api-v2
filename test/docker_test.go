package test

import (
	"context"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/Capstone-auto-grader/grader-api-v2/internal/graderd"
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
			d := graderd.NewDockerClient(test.host, test.version)
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
		tasks     []*graderd.Task
		err       error
	}{
		{
			desc:      "one task",
			host:      "http://localhost:2376",
			version:   "1.40",
			imageName: "helloworld",
			tasks: []*graderd.Task{
				{
					AssignmentID: "helloworld",
					StudentName:  "some name",
					Urn:          "some:urn:key",
					ZipKey:       "some:zip:key",
				},
			},
			err: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			d := graderd.NewDockerClient(test.host, test.version)
			db := graderd.NewMockDatabase()
			ctx := context.Background()
			// create tasks
			ids, err := d.CreateTasks(ctx, test.tasks, db)
			if err != nil {
				if want, got := test.err, err; !reflect.DeepEqual(want, got) {
					t.Errorf("unexpected error:\n- want: %+v\n- got: %+v\n", want, got)
				}
				return
			}
			if want, got := len(test.tasks), len(ids); want != got {
				t.Errorf("unexpected number of IDs:\n- want: %+v\n- got: %+v\n", want, got)
			}
			t.Logf("\n- Container IDs: %+v", ids)
		})
	}
}

func TestDockerClient_CreateAndStartTasks(t *testing.T) {
	tests := []struct {
		desc      string
		host      string
		version   string
		imageName string
		tasks     []*graderd.Task
		err       error
	}{
		{
			desc:      "one task, foreign image",
			host:      "http://localhost:2376",
			version:   "1.40",
			imageName: "hello-world",
			tasks: []*graderd.Task{
				{
					AssignmentID: "hello-world",
					StudentName:  "some name",
					Urn:          "some:urn:key",
					ZipKey:       "some:zip:key",
				},
			},
			err: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			d := graderd.NewDockerClient(test.host, test.version)
			db := graderd.NewMockDatabase()
			ctx := context.Background()
			// create tasks
			ids, err := d.CreateTasks(ctx, test.tasks, db)
			if err != nil {
				if want, got := test.err, err; !reflect.DeepEqual(want, got) {
					t.Errorf("unexpected error:\n- want: %+v\n- got: %+v\n", want, got)
				}
				return
			}
			if want, got := len(test.tasks), len(ids); want != got {
				t.Errorf("unexpected number of IDs:\n- want: %+v\n- got: %+v\n", want, got)
			}
			t.Logf("\n- Container IDs: %+v", ids)
			// start tasks
			err = d.StartTasks(ctx, ids, db)
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
