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

var RevokeRefreshFlags []cli.Flag = []cli.Flag{
	cli.StringFlag{
		Name:  "grpc-url, g",
		Value: "127.0.0.1:5557",
		Usage: "gRPC host and port",
	},
}

func RevokeRefresh(c *cli.Context) error {
	ctx := context.Background()

	grpc, err := setup.SetupGrpcClient(ctx, c)
	if err != nil {
		return err
	}

	clientID := c.Args().Get(0)
	if clientID == "" {
		return errors.New("client id must be provided as first argument")
	}

	userID := c.Args().Get(1)
	if userID == "" {
		return errors.New("user id must be provided as first argument")
	}

	resp, err := grpc.RevokeRefresh(ctx, &api.RevokeRefreshReq{
		UserId:   userID,
		ClientId: clientID,
	})
	if err != nil {
		return err
	}

	if resp.NotFound {
		return errors.New("client/user not found")
	}

	p, err := utils.PrettifyJSON(resp)
	if err != nil {
		return err
	}

	fmt.Println(p)

	return nil
}
