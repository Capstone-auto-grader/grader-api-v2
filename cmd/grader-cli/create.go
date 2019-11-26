package main

import (
	"context"
	"io/ioutil"

	"github.com/urfave/cli"

	pb "github.com/Capstone-auto-grader/grader-api-v2/graderpb"
)

// CreateAssignment builds a cli.Command for calling the CreateAssignment endpoint.
func CreateAssignment() cli.Command {
	command := cli.Command{
		Name:        "create",
		Description: "create creates an assignment based on a dockerfile and run script",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "name, n",
				Usage: "image name",
			},
			cli.StringFlag{
				Name:  "file, f",
				Usage: "image file",
			},
		},
		Action: CreateAssignmentAction(),
	}

	return command
}

// CreateAssignmentAction builds the cli.ActionFunc for actually calling the endpoint with the gRPC client.
func CreateAssignmentAction() cli.ActionFunc {
	return func(c *cli.Context) error {
		ctx := context.Background()
		addr := c.GlobalString("addr")
		cert := c.GlobalString("cert")

		name := c.String("name")
		file := c.String("file")
		// open tarball
		t, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}

		req := &pb.CreateAssignmentRequest{
			ImageName: name,
			ImageTar:  t,
		}

		client := NewClient(cert, addr)
		if _, err := client.CreateAssignment(ctx, req); err != nil {
			return err
		}

		return nil
	}
}
