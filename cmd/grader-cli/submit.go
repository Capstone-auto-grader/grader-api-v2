package main

import (
	"context"
	"fmt"

	"github.com/urfave/cli"

	pb "github.com/Capstone-auto-grader/grader-api-v2/graderpb"
)

// SubmitForGrading builds a cli.Command for calling the SubmitForGrading endpoint.
func SubmitForGrading() cli.Command {
	command := cli.Command{
		Name:        "submit",
		Description: "submit sends an assignment to be graded to the grader API",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "id, i",
				Usage: "assignment id",
			},
			cli.StringFlag{
				Name:  "urn, u",
				Usage: "assignment urn",
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
		Action: SubmitForGradingAction(),
	}

	return command
}

// SubmitForGradingAction builds the cli.ActionFunc for actually calling the endpoint with the gRPC client.
func SubmitForGradingAction() cli.ActionFunc {
	return func(c *cli.Context) error {
		ctx := context.Background()
		addr := c.GlobalString("addr")
		cert := c.GlobalString("cert")

		id := c.String("id")
		urn := c.String("urn")
		zipkey := c.String("zipkey")
		name := c.String("name")

		req := &pb.SubmitForGradingRequest{
			Tasks: []*pb.Task{
				{
					AssignmentId: id,
					UrnKey:       urn,
					ZipKey:       zipkey,
					StudentName:  name,
					Timeout:      30,
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
