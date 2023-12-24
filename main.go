package main

import (
	"errors"
	"os"

	"github.com/twoojoo/dexcli/signin"
	"github.com/urfave/cli"
)

var (
	errCommandNotAvailable   = errors.New("command not available")
	usageCommandNotAvaliable = "command not available"
)

func main() {
	err := runCLI()

	if err != nil {
		panic(err)
	}
}

func runCLI() error {
	app := cli.NewApp()
	app.Name = "dexcli"
	app.Usage = "a Command Line Interface for Dex"

	app.Commands = []cli.Command{
		{
			Name:   "signin",
			Usage:  "Performs a sign-in using a browser",
			Flags:  signin.Flags,
			Action: signin.Signin,
		},
		{
			Name:  "connector",
			Usage: "connector-related commands",
			Subcommands: []cli.Command{
				{
					Name:   "list",
					Usage:  usageCommandNotAvaliable,
					Action: func(c *cli.Context) error { return errCommandNotAvailable },
				},
				{
					Name:   "create",
					Usage:  usageCommandNotAvaliable,
					Action: func(c *cli.Context) error { return errCommandNotAvailable },
				},
				{
					Name:   "delete",
					Usage:  usageCommandNotAvaliable,
					Action: func(c *cli.Context) error { return errCommandNotAvailable },
				},
				{
					Name:   "update",
					Usage:  usageCommandNotAvaliable,
					Action: func(c *cli.Context) error { return errCommandNotAvailable },
				},
			},
		},
		{
			Name:  "client",
			Usage: "client-related commands",
			Subcommands: []cli.Command{
				{
					Name:   "list",
					Usage:  usageCommandNotAvaliable,
					Action: func(c *cli.Context) error { return errCommandNotAvailable },
				},
				{
					Name:   "create",
					Usage:  usageCommandNotAvaliable,
					Action: func(c *cli.Context) error { return errCommandNotAvailable },
				},
				{
					Name:   "delete",
					Usage:  usageCommandNotAvaliable,
					Action: func(c *cli.Context) error { return errCommandNotAvailable },
				},
				{
					Name:   "update",
					Usage:  usageCommandNotAvaliable,
					Action: func(c *cli.Context) error { return errCommandNotAvailable },
				},
			},
		},
		{
			Name:  "password",
			Usage: "password-related commands",
			Subcommands: []cli.Command{
				{
					Name:   "list",
					Usage:  usageCommandNotAvaliable,
					Action: func(c *cli.Context) error { return errCommandNotAvailable },
				},
				{
					Name:   "create",
					Usage:  usageCommandNotAvaliable,
					Action: func(c *cli.Context) error { return errCommandNotAvailable },
				},
				{
					Name:   "delete",
					Usage:  usageCommandNotAvaliable,
					Action: func(c *cli.Context) error { return errCommandNotAvailable },
				},
				{
					Name:   "update",
					Usage:  usageCommandNotAvaliable,
					Action: func(c *cli.Context) error { return errCommandNotAvailable },
				},
			},
		},
	}

	return app.Run(os.Args)
}
