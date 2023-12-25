package connector

import (
	// "context"

	// "github.com/twoojoo/dexctl/setup"
	"github.com/urfave/cli"
)

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

// func CreateConnector(ctx context.Context, c *cli.Context) error {
// 	grpc, err := setup.SetupGrpcClient(ctx, c)
// 	if err != nil {
// 		return err
// 	}

// 	grpc.

// 	return nil
// }
