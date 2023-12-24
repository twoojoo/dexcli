package signin

import (
	"github.com/urfave/cli"
)

var Flags []cli.Flag = []cli.Flag{
	cli.StringFlag{
		Name:  "dex-base-url",
		Value: "http://127.0.0.1:5556",
	},
	cli.UintFlag{
		Name:  "port, p",
		Value: 3000,
	},
	cli.StringFlag{
		Name:  "client-id, cid",
		Value: "example-app",
	},
	cli.StringFlag{
		Name:  "client-secret, csec",
		Value: "example-app-secret",
	},
	cli.StringFlag{
		Name:  "state, t",
		Value: "",
	},
	cli.StringSliceFlag{
		Name:  "scopes, s",
		Value: &cli.StringSlice{},
	},
	cli.StringFlag{
		Name:     "browser, b",
		Required: false,
	},
}
