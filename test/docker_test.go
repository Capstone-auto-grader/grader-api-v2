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
			err := d.CreateAssignment(ctx, test.imageName, test.imageTar)
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
