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

var GetPasswordFlags []cli.Flag = []cli.Flag{
	cli.StringFlag{
		Name:  "grpc-url, g",
		Value: "127.0.0.1:5557",
		Usage: "gRPC host and port",
	},
}

func GetPassword(c *cli.Context) error {
	ctx := context.Background()

	email, err := utils.ParseEmail(c.Args().Get(0))
	if err != nil {
		return err
	}

	grpc, err := setup.SetupGrpcClient(ctx, c)
	if err != nil {
		return err
	}

	resp, err := grpc.ListPasswords(ctx, &api.ListPasswordReq{})
	if err != nil {
		return err
	}

	var password *api.Password = nil
	for _, p := range resp.Passwords {
		if p.Email == email {
			password = p
		}
	}

	if password == nil {
		return errors.New("password not found")
	}

	p, err := utils.PrettifyJSON(password)
	if err != nil {
		return err
	}

	fmt.Println(p)

	return nil
}
