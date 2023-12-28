package password

import (
	"context"
	"errors"
	"fmt"
	"net/mail"

	"github.com/dexidp/dex/api/v2"
	"github.com/twoojoo/dexctl/setup"
	"github.com/twoojoo/dexctl/utils"
	"github.com/urfave/cli"
)

var DeleteConnectorFlags []cli.Flag = []cli.Flag{
	cli.StringFlag{
		Name:  "grpc-url, g",
		Value: "127.0.0.1:5557",
		Usage: "gRPC host and port",
	},
}

func DeleteConnector(c *cli.Context) error {
	ctx := context.Background()

	grpc, err := setup.SetupGrpcClient(ctx, c)
	if err != nil {
		return err
	}

	id := c.Args().Get(0)
	if id == "" {
		return errors.New("client id must be provided as first argument")
	}

	resp, err := grpc.DeleteConnector(ctx, &api.DeleteConnectorReq{
		ID: id,
	})

	if err != nil {
		return err
	}

	if resp.NotFound {
		return errors.New("password not found")
	}

	p, err := utils.PrettifyJSON(resp)
	if err != nil {
		return err
	}

	fmt.Println(p)

	return nil
}
