package connector

import "github.com/urfave/cli"

var CreateConnectorFlags []cli.Flag = []cli.Flag{
	cli.StringFlag{
		Name:  "grpc-url, g",
		Value: "127.0.0.1:5557",
		Usage: "GRPC host and port",
	},
	cli.StringFlag{
		Name:  "id, i",
		Value: "random UUID",
	},
}
