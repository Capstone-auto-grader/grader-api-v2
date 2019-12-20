// Grader-cli is a client for grader API provided by grader.pb.go.
//
// Usage
//
//	go build -o *.go grader-cli
//	./grader-cli
//
// This client currently supports the following endpoints:
//	SubmitForGrading()
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
		Submit(),
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
