package refresh

import (
	"context"
	"errors"
	"fmt"

	"github.com/dexidp/dex/api/v2"
	"github.com/twoojoo/dexctl/setup"
	"github.com/twoojoo/dexctl/utils"
	"github.com/urfave/cli"
)

var ListRefreshFlags []cli.Flag = []cli.Flag{
	cli.StringFlag{
		Name:  "grpc-url, g",
		Value: "127.0.0.1:5557",
		Usage: "gRPC host and port",
	},
}

func ListRefresh(c *cli.Context) error {
	ctx := context.Background()

	grpc, err := setup.SetupGrpcClient(ctx, c)
	if err != nil {
		return err
	}

	id := c.Args().Get(0)
	if id == "" {
		return errors.New("user id must be provided as first argument")
	}

	resp, err := grpc.ListRefresh(ctx, &api.ListRefreshReq{UserId: id})
	if err != nil {
		return err
	}

	p, err := utils.PrettifyJSON(resp.RefreshTokens)
	if err != nil {
		return err
	}

	fmt.Println(p)

	return nil
}
