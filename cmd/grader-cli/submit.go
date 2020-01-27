package main

import (
	"context"
	"fmt"
	"github.com/urfave/cli"

	pb "github.com/Capstone-auto-grader/grader-api-v2/graderpb"
)

// Submit builds a cli.Command for calling the SubmitForGrading endpoint.
func Submit() cli.Command {
	command := cli.Command{
		Name:        "submit",
		Description: "submit sends an assignment to be graded to the grader API",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "image_name, i",
				Usage: "image name",
			},
			cli.StringFlag{
				Name:  "uri, u",
				Usage: "assignment uri",
			},
			cli.StringFlag{
				Name:  "zipkey, z",
				Usage: "zip key (s3 key)",
			},
			cli.StringFlag{
				Name:  "name, n",
				Usage: "student's name",
			},
		},
		Action: SubmitAction(),
	}

	return command
}

// SubmitAction builds the cli.ActionFunc for actually calling the endpoint with the gRPC client.
func SubmitAction() cli.ActionFunc {
	return func(c *cli.Context) error {
		ctx := context.Background()
		addr := c.GlobalString("addr")
		cert := c.GlobalString("cert")

		id := c.String("image_name")
		uri := c.String("uri")
		zipkey := c.String("zipkey")
		name := c.String("name")

		req := &pb.SubmitForGradingRequest{
			Tasks: []*pb.Task{
				{
					ImageName: id,
					TestKey:       uri,
					ZipKey:       zipkey,
					StudentName:  name,
				},
			},
		}

		client := NewClient(cert, addr)
		resp, err := client.SubmitForGrading(ctx, req)
		if err != nil {
			return err
		}

		fmt.Println(resp.String())
		return nil
	}
}
