package client

import (
	"context"
	"errors"
	"fmt"

	"github.com/dexidp/dex/api/v2"
	"github.com/twoojoo/dexctl/setup"
	"github.com/twoojoo/dexctl/utils"
	"github.com/urfave/cli"
)

var UpdateClientFlags []cli.Flag = []cli.Flag{
	cli.StringFlag{
		Name:  "grpc-url, g",
		Value: "127.0.0.1:5557",
		Usage: "gRPC host and port",
	},
	cli.StringFlag{
		Name:     "name, n",
		Required: false,
	},
	cli.StringSliceFlag{
		Name:     "redirect-uris, r",
		Required: false,
	},
	cli.StringFlag{
		Name:     "logo-url, l",
		Required: false,
	},
	cli.StringSliceFlag{
		Name:     "trusted-peers, t",
		Required: false,
	},
}

func UpdateClient(c *cli.Context) error {
	ctx := context.Background()

	grpc, err := setup.SetupGrpcClient(ctx, c)
	if err != nil {
		return err
	}

	id := c.Args().Get(0)
	if id == "" {
		return errors.New("client id must be provided as first argument")
	}

	resp, err := grpc.UpdateClient(ctx, &api.UpdateClientReq{
		Id:           id,
		RedirectUris: c.StringSlice("redirect-uris"),
		TrustedPeers: c.StringSlice("trusted-peers"),
		Name:         c.String("name"),
		LogoUrl:      c.String("logo-url"),
	})

	if err != nil {
		return err
	}

	p, err := utils.PrettifyJSON(resp)
	if err != nil {
		return err
	}

	fmt.Println(p)

	return nil
}
