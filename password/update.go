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

var UpdatePasswordFlags []cli.Flag = []cli.Flag{
	cli.StringFlag{
		Name:  "grpc-url, g",
		Value: "127.0.0.1:5557",
		Usage: "gRPC host and port",
	},
	cli.StringFlag{
		Name: "username, u",
	},
	cli.StringFlag{
		Name:  "hash, s", //conflict with help on using h as alias
		Usage: "bcrypt hash of the password: $(echo <password> | htpasswd -BinC 10 admin | cut -d: -f2)",
		Value: "$2a$10$2b2cU8CPhOTaGrs1HRQuAueS7JTT5ZHsHSzYiFPm1leZck7Mc8T4W", //"password"
	},
}

func UpdatePassword(c *cli.Context) error {
	ctx := context.Background()

	email, err := utils.ParseEmail(c.Args().Get(0))
	if err != nil {
		return err
	}

	grpc, err := setup.SetupGrpcClient(ctx, c)
	if err != nil {
		return err
	}

	resp, err := grpc.UpdatePassword(ctx, &api.UpdatePasswordReq{
		Email:       email,
		NewHash:     []byte(c.String("hash")),
		NewUsername: c.String("username"),
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
