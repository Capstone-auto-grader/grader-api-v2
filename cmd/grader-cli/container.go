package main

import (
	"context"
	"fmt"
	pb "github.com/Capstone-auto-grader/grader-api-v2/graderpb"
	"github.com/urfave/cli"
	"io/ioutil"
)

func MakeContainer() cli.Command {
	command := cli.Command{
		Name: "build_container",
		Description: "Given a TAR file, builds a container",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name: "file_path, f",
				Usage: "file path",
			},
			cli.StringFlag{
				Name: "container_tag, t",
				Usage: "container tag",
			},
		},
		Action: MakeContainerAction(),
	}
	return command
}

func MakeContainerAction() cli.ActionFunc {
	return func(c *cli.Context) error {
		ctx := context.Background()
		addr := c.GlobalString("addr")
		cert := c.GlobalString("cert")

		filepath := c.String("file_path")
		tag := c.String("container_tag")
		fmt.Println("TAG ", tag)
		tar, err := ioutil.ReadFile(filepath)
		if err != nil {
			return fmt.Errorf("error reading file %s: %s", filepath, err)

		}
		req := &pb.CreateImageRequest{
			ImageName:tag,
			ImageTar: tar}

		client := NewClient(cert, addr)
		resp, err := client.CreateImage(ctx, req)
		if err != nil {
			return err
		}
		fmt.Println(resp.String())
		return nil
	}
}
