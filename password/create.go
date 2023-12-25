package password

import (
	"context"
	"fmt"
	"net/mail"

	"github.com/dexidp/dex/api/v2"
	"github.com/google/uuid"
	"github.com/twoojoo/dexctl/setup"
	"github.com/twoojoo/dexctl/utils"
	"github.com/urfave/cli"
)

var CreatePasswordFlags []cli.Flag = []cli.Flag{
	cli.StringFlag{
		Name:  "grpc-url, g",
		Value: "127.0.0.1:5557",
		Usage: "gRPC host and port",
	},
	cli.StringFlag{
		Name:     "username, u",
		Required: true,
	},
	cli.StringFlag{
		Name:  "hash, s", //conflict with help on using h as alias
		Usage: "bcrypt hash of the password: $(echo <password> | htpasswd -BinC 10 admin | cut -d: -f2)",
		Value: "$2a$10$2b2cU8CPhOTaGrs1HRQuAueS7JTT5ZHsHSzYiFPm1leZck7Mc8T4W",
	},
	cli.StringFlag{
		Name:  "id, i",
		Value: "random UUID",
	},
}

func CreatePassword(c *cli.Context) error {
	ctx := context.Background()

	grpc, err := setup.SetupGrpcClient(ctx, c)
	if err != nil {
		return err
	}

	id := c.String("id")
	if id == "random UUID" || id == "" {
		uid, err := uuid.NewUUID()
		if err != nil {
			return err
		}

		id = uid.String()
	}

	email := c.Args().Get(0)
	_, err = mail.ParseAddress(email)
	if err != nil {
		return err
	}

	pwd := &api.Password{
		Email:    email,
		Hash:     []byte(c.String("hash")),
		Username: c.String("username"),
		UserId:   id,
	}

	resp, err := grpc.CreatePassword(ctx, &api.CreatePasswordReq{Password: pwd})
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
