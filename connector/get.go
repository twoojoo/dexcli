package password

import (
	"context"
	"errors"
	"fmt"

	"github.com/dexidp/dex/api/v2"
	"github.com/twoojoo/dexctl/setup"
	"github.com/twoojoo/dexctl/utils"
	"github.com/urfave/cli"
)

var GetPasswordFlags []cli.Flag = []cli.Flag{
	cli.StringFlag{
		Name:  "grpc-url, g",
		Value: "127.0.0.1:5557",
		Usage: "gRPC host and port",
	},
}

func GetPassword(c *cli.Context) error {
	ctx := context.Background()

	id := c.Args().Get(0)
	if id == "" {
		return errors.New("client id must be provided as first argument")
	}

	grpc, err := setup.SetupGrpcClient(ctx, c)
	if err != nil {
		return err
	}

	resp, err := grpc.ListConnectors(ctx, &api.ListConnectorReq{})

	if err != nil {
		return err
	}

	var connector *api.Connector = nil
	for _, c := range resp.Connectors {
		if p.ID == id {
			connector = c
		}
	}

	if connector == nil {
		return errors.New("password not found")
	}

	p, err := utils.PrettifyJSON(connector)
	if err != nil {
		return err
	}

	fmt.Println(p)

	return nil
}
