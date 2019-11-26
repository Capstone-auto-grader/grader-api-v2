// grader-cli is a wrapper around the GraderClient gRPC client provided by graderpb.
//
// Usage
//
// 		go build -o *.go grader-cli
//		./grader-cli
//
// This file provides all the definition of flags and commands for the grader-cli.
package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

// main is the entry-point for running the grader-cli.
func main() {
	// create cli app
	app := cli.NewApp()
	app.Name = "grader-cli"
	app.Usage = "a cli client for grader API"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "addr, a",
			Usage: "address of the endpoint",
		},
		cli.StringFlag{
			Name:  "cert, c",
			Usage: "location of the public cert",
		},
	}
	app.Commands = cli.Commands{
		SubmitForGrading(),
		CreateAssignment(),
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
