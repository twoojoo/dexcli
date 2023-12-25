package client

import (
	"context"
	"fmt"

	"github.com/dexidp/dex/api/v2"
	"github.com/google/uuid"
	"github.com/twoojoo/dexctl/setup"
	"github.com/twoojoo/dexctl/utils"
	"github.com/urfave/cli"
)

var CreateClientFlags []cli.Flag = []cli.Flag{
	cli.StringFlag{
		Name:  "grpc-url, g",
		Value: "127.0.0.1:5557",
		Usage: "gRPC host and port",
	},
	cli.StringFlag{
		Name:     "name, n",
		Required: true,
	},
	cli.StringFlag{
		Name:  "client-id, cid",
		Value: "random UUID",
	},
	cli.StringFlag{
		Name:  "client-secret, csec",
		Value: "random string",
	},
	cli.StringSliceFlag{
		Name:  "redirect-uris, r",
		Value: &cli.StringSlice{"http://127.0.0.1:3000/callback"},
	},
	cli.StringFlag{
		Name:     "logo-url, l",
		Required: false,
	},
	cli.StringSliceFlag{
		Name:     "trusted-peers, t",
		Required: false,
	},
	cli.BoolFlag{
		Name:     "public, p",
		Required: false,
	},
}

func CreateClient(c *cli.Context) error {
	ctx := context.Background()

	grpc, err := setup.SetupGrpcClient(ctx, c)
	if err != nil {
		return err
	}

	clientId := c.String("client-id")
	if clientId == "random UUID" || clientId == "" {
		uid, err := uuid.NewUUID()
		if err != nil {
			return err
		}

		clientId = uid.String()
	}

	clientSecret := c.String("client-secret")
	if clientSecret == "random string" || clientSecret == "" {
		clientSecret = utils.RandomString(utils.LettersSet, 15)
	}

	client := &api.Client{
		Id:           clientId,
		Secret:       clientSecret,
		Name:         c.String("name"),
		RedirectUris: c.StringSlice("redirect-uris"),
		LogoUrl:      c.String("logo-url"),
		TrustedPeers: c.StringSlice("trusted-peers"),
		Public:       c.Bool("public"),
	}

	resp, err := grpc.CreateClient(ctx, &api.CreateClientReq{Client: client})
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
