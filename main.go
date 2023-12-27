package main

import (
	"errors"
	"log"
	"os"

	"github.com/twoojoo/dexctl/base"
	"github.com/twoojoo/dexctl/client"
	"github.com/twoojoo/dexctl/password"
	"github.com/twoojoo/dexctl/refresh"
	"github.com/twoojoo/dexctl/token"
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
			Name:   "signon",
			Usage:  "Performs a sign-on (opens up a browser page)",
			Flags:  base.SignonFlags,
			Action: base.Signon,
		},
		{
			Name:   "version",
			Usage:  "Get the version of the Dex server",
			Flags:  base.VersionFlags,
			Action: base.Version,
		},
		{
			Name:  "token",
			Usage: "token-related commands",
			Subcommands: []cli.Command{
				{
					Name:   "verify",
					Usage:  "Attempts to verify an access token or ID token returning the token claims",
					Flags:  token.VerifyTokenFlags,
					Action: token.VerifyToken,
				},
			},
		},
		// {
		// 	Name:  "connector",
		// 	Usage: "Connector-related commands",
		// 	Subcommands: []cli.Command{
		// 		{
		// 			Name:   "list",
		// 			Usage:  usageCommandNotAvaliable,
		// 			Action: func(c *cli.Context) error { return errCommandNotAvailable },
		// 		},
		// 		{
		// 			Name:   "create",
		// 			Usage:  usageCommandNotAvaliable,
		// 			Action: func(c *cli.Context) error { return errCommandNotAvailable },
		// 		},
		// 		{
		// 			Name:   "delete",
		// 			Usage:  usageCommandNotAvaliable,
		// 			Action: func(c *cli.Context) error { return errCommandNotAvailable },
		// 		},
		// 		{
		// 			Name:   "update",
		// 			Usage:  usageCommandNotAvaliable,
		// 			Action: func(c *cli.Context) error { return errCommandNotAvailable },
		// 		},
		// 	},
		// },
		{
			Name:  "client",
			Usage: "Client-related commands",
			Subcommands: []cli.Command{
				{
					Name:   "list",
					Usage:  usageCommandNotAvaliable,
					Action: func(c *cli.Context) error { return errCommandNotAvailable },
				},
				{
					Name:   "create",
					Usage:  "Add a new client",
					Flags:  client.CreateClientFlags,
					Action: client.CreateClient,
				},
				{
					Name:      "delete",
					Usage:     "Delete an existing client",
					ArgsUsage: "<client_id>",
					Flags:     client.DeleteClientFlags,
					Action:    client.DeleteClient,
				},
				{
					Name:      "update",
					Usage:     "Update an existing client",
					ArgsUsage: "<client_id>",
					Flags:     client.UpdateClientFlags,
					Action:    client.UpdateClient,
				},
			},
		},
		{
			Name:  "password",
			Usage: "Password-related commands",
			Subcommands: []cli.Command{
				{
					Name:      "get",
					Usage:     "Retrieves a password (user)",
					ArgsUsage: "<email>",
					Flags:     password.GetPasswordFlags,
					Action:    password.GetPassword,
				},
				{
					Name:   "list",
					Usage:  "Lists all passwords (users)",
					Flags:  password.ListPasswordFlags,
					Action: password.ListPassword,
				},
				{
					Name:      "create",
					Usage:     "Add a new password (user)",
					ArgsUsage: "<email>",
					Flags:     password.CreatePasswordFlags,
					Action:    password.CreatePassword,
				},
				{
					Name:      "delete",
					Usage:     "Delete an existing password (user)",
					ArgsUsage: "<email>",
					Flags:     password.DeletePasswordFlags,
					Action:    password.DeletePassword,
				},
				{
					Name:      "update",
					Usage:     "Update an existing password (user)",
					ArgsUsage: "<email>",
					Flags:     password.UpdatePasswordFlags,
					Action:    password.UpdatePassword,
				},
				{
					Name:      "verify",
					Usage:     "Attempts to verify a password for a user email",
					ArgsUsage: "<email>",
					Flags:     password.VerifyPasswordFlags,
					Action:    password.VerifyPassword,
				},
			},
		},
		{
			Name:  "refresh",
			Usage: "Refresh-related commands",
			Subcommands: []cli.Command{
				{
					Name:      "list",
					ArgsUsage: "<user_id>",
					Flags:     refresh.ListRefreshFlags,
					Action:    refresh.ListRefresh,
				},
				{
					Name:      "revoke",
					ArgsUsage: "<client_id> <user_id>",
					Flags:     refresh.RevokeRefreshFlags,
					Action:    refresh.RevokeRefresh,
				},
			},
		},
	}

	return app.Run(os.Args)
}
