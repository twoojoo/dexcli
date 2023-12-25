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

var VerifyPasswordFlags []cli.Flag = []cli.Flag{
	cli.StringFlag{
		Name:  "grpc-url, g",
		Value: "127.0.0.1:5557",
		Usage: "gRPC host and port",
	},
}

func VerifyPassword(c *cli.Context) error {
	ctx := context.Background()

	grpc, err := setup.SetupGrpcClient(ctx, c)
	if err != nil {
		return err
	}

	email := c.Args().Get(0)
	_, err = mail.ParseAddress(email)
	if err != nil {
		return err
	}

	password := c.Args().Get(1)
	if password == "" {
		return errors.New("password must be provided as second argument")
	}

	resp, err := grpc.VerifyPassword(ctx, &api.VerifyPasswordReq{
		Email:    email,
		Password: password,
	})

	if err != nil {
		return err
	}

	if resp.NotFound {
		return errors.New("password not found")
	}

	if !resp.Verified {
		return errors.New("failed to verify password")
	}

	p, err := utils.PrettifyJSON(resp)
	if err != nil {
		return err
	}

	fmt.Println(p)

	return nil
}
