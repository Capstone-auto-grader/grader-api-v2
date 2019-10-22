package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

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
