package main

import (
	"errors"
	"log"
	"os"

	"github.com/twoojoo/dexcli/signin"
	"github.com/urfave/cli"
)

var (
	errCommandNotAvailable = errors.New("command not available")
)

func main() {
	app := cli.NewApp()
	app.Name = "dexcli"
	app.Usage = "a CLI for Dex"

	app.Commands = []cli.Command{
		{
			Name:   "signin",
			Flags:  signin.Flags,
			Action: signin.Signin,
		},
		{
			Name: "connector",
			Subcommands: []cli.Command{
				{
					Name:   "list",
					Action: func(c *cli.Context) error { return errCommandNotAvailable },
				},
				{
					Name:   "create",
					Action: func(c *cli.Context) error { return errCommandNotAvailable },
				},
				{
					Name:   "delete",
					Action: func(c *cli.Context) error { return errCommandNotAvailable },
				},
				{
					Name:   "update",
					Action: func(c *cli.Context) error { return errCommandNotAvailable },
				},
			},
		},
		{
			Name: "client",
			Subcommands: []cli.Command{
				{
					Name:   "list",
					Action: func(c *cli.Context) error { return errCommandNotAvailable },
				},
				{
					Name:   "create",
					Action: func(c *cli.Context) error { return errCommandNotAvailable },
				},
				{
					Name:   "delete",
					Action: func(c *cli.Context) error { return errCommandNotAvailable },
				},
				{
					Name:   "update",
					Action: func(c *cli.Context) error { return errCommandNotAvailable },
				},
			},
		},
		{
			Name: "password",
			Subcommands: []cli.Command{
				{
					Name:   "list",
					Action: func(c *cli.Context) error { return errCommandNotAvailable },
				},
				{
					Name:   "create",
					Action: func(c *cli.Context) error { return errCommandNotAvailable },
				},
				{
					Name:   "delete",
					Action: func(c *cli.Context) error { return errCommandNotAvailable },
				},
				{
					Name:   "update",
					Action: func(c *cli.Context) error { return errCommandNotAvailable },
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
