package main

import (
	"fmt"
	"os"

	"github.com/cecchisandrone/go-examples/functions"
	"github.com/cecchisandrone/go-examples/maps"
	"github.com/cecchisandrone/go-examples/slices"
	"github.com/cecchisandrone/go-examples/structs"

	"github.com/urfave/cli"
)

type example interface {
	Start()
}

func main() {

	examples := map[string]example{
		"functions": functions.Functions{},
		"slices":    slices.Slices{},
		"maps":      maps.Maps{},
		"structs":   structs.Structs{},
	}

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

		example := examples[topic]

		if example != nil {
			example.Start()
		} else {
			fmt.Println("Bad example name")
			os.Exit(1)
		}

		return nil
	}
	app.Run(os.Args)
}
