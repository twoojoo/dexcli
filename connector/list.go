package password

import (
	"context"
	"fmt"

	"github.com/dexidp/dex/api/v2"
	"github.com/twoojoo/dexctl/setup"
	"github.com/twoojoo/dexctl/utils"
	"github.com/urfave/cli"
)

var ListConnectorFlags []cli.Flag = []cli.Flag{
	cli.StringFlag{
		Name:  "grpc-url, g",
		Value: "127.0.0.1:5557",
		Usage: "gRPC host and port",
	},
}

func ListConnectors(c *cli.Context) error {
	ctx := context.Background()

	grpc, err := setup.SetupGrpcClient(ctx, c)
	if err != nil {
		return err
	}

	resp, err := grpc.ListConnectors(ctx, &api.ListConnectorReq{})

	if err != nil {
		return err
	}

	p, err := utils.PrettifyJSON(resp.Passwords)
	if err != nil {
		return err
	}

	fmt.Println(p)

	return nil
}
