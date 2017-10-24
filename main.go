package main

import (
	"os"

	"github.com/cecchisandrone/go-examples/maps"
	"github.com/cecchisandrone/go-examples/slices"
	"github.com/cecchisandrone/go-examples/structs"

	"github.com/urfave/cli"
)

func main() {

	var topic string
	app := cli.NewApp()
	app.Version = "1.0"
	app.Name = "Go Examples"
	app.Description = "A set of examples in Go"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "topic, t",
			Usage:       "Run examples on specified topic",
			Destination: &topic,
		},
	}
	app.Action = func(c *cli.Context) error {

		if c.NumFlags() == 0 && c.NArg() == 0 {
			cli.ShowAppHelpAndExit(c, 0)
		}

		if topic == "structs" {
			structs.Start()
		} else if topic == "slices" {
			slices.Start()
		} else if topic == "maps" {
			maps.Start()
		}

		return nil
	}
	app.Run(os.Args)
}
