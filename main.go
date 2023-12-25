package main

import (
	"errors"
	"log"
	"os"

	"github.com/twoojoo/dexctl/base"
	"github.com/twoojoo/dexctl/client"
	"github.com/twoojoo/dexctl/connector"
	"github.com/twoojoo/dexctl/password"
	"github.com/urfave/cli"
)

var (
	errCommandNotAvailable   = errors.New("command not available")
	usageCommandNotAvaliable = "command not available"
)

func main() {
	err := runCLI()

	if err != nil {
		log.Fatal(err)
	}
}

func runCLI() error {
	app := cli.NewApp()
	app.Name = "dexctl"
	app.EnableBashCompletion = true
	app.Usage = "a Command Line Interface for Dex"

	app.Commands = []cli.Command{
		{
			Name:   "signin",
			Usage:  "Performs a sign-in using a browser",
			Flags:  base.SigninFlags,
			Action: base.Signin,
		},
		{
			Name:   "verify",
			Usage:  "attempts to verify an access_token",
			Flags:  base.VerifyFlags,
			Action: base.Verify,
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
					Flags:  connector.CreateConnectorFlags,
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
					Flags:  client.CreateClientFlags,
					Action: client.CreateClient,
				},
				{
					Name:   "delete",
					Usage:  usageCommandNotAvaliable,
					Action: func(c *cli.Context) error { return errCommandNotAvailable },
				},
				{
					Name:      "update",
					ArgsUsage: "<client_id>",
					Flags:     client.UpdateClientFlags,
					Action:    client.UpdateClient,
				},
			},
		},
		{
			Name:  "password",
			Usage: "password-related commands",
			Subcommands: []cli.Command{
				{
					Name:      "get",
					ArgsUsage: "<email>",
					Flags:     password.GetPasswordFlags,
					Action:    password.GetPassword,
				},
				{
					Name:   "list",
					Flags:  password.ListPasswordFlags,
					Action: password.ListPassword,
				},
				{
					Name:      "create",
					ArgsUsage: "<email>",
					Flags:     password.CreatePasswordFlags,
					Action:    password.CreatePassword,
				},
				{
					Name:      "delete",
					ArgsUsage: "<email>",
					Flags:     password.DeletePasswordFlags,
					Action:    password.DeletePassword,
				},
				{
					Name:      "update",
					ArgsUsage: "<email>",
					Flags:     password.UpdatePasswordFlags,
					Action:    password.UpdatePassword,
				},
				{
					Name:      "verify",
					ArgsUsage: "<email>",
					Flags:     password.VerifyPasswordFlags,
					Action:    password.VerifyPassword,
				},
			},
		},
	}

	return app.Run(os.Args)
}
