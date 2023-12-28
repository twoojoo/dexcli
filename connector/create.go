package password

import (
	"context"
	"errors"
	"fmt"

	"github.com/dexidp/dex/api/v2"
	"github.com/google/uuid"
	"github.com/twoojoo/dexctl/setup"
	"github.com/twoojoo/dexctl/utils"
	"github.com/urfave/cli"
)

var CreateConnectorFlags []cli.Flag = []cli.Flag{
	cli.StringFlag{
		Name:  "grpc-url, g",
		Value: "127.0.0.1:5557",
		Usage: "gRPC host and port",
	},
	cli.StringFlag{
		Name:  "id, i",
		Value: "random UUID",
	},
	cli.StringFlag{
		Name:     "name, n",
		Required: true,
	},
	cli.StringFlag{
		Name:     "type, t",
		Required: true,
	},
	cli.StringFlag{
		Name:  "config, c",
		Value: "{}",
	},
}

func CreateConnector(c *cli.Context) error {
	ctx := context.Background()

	id, err := utils.ParseRandomUUID(c.String("id"))
	if err != nil {
		return err
	}

	grpc, err := setup.SetupGrpcClient(ctx, c)
	if err != nil {
		return err
	}

	resp, err := grpc.CreateConnector(ctx, &api.CreateConnectorReq{
		Connector: &api.Connector{
			ID:     id,
			Type:   c.String("type"),
			Name:   c.String("name"),
			Config: []byte(c.String("config")),
		},
	})
	if err != nil {
		return err
	}

	if resp.AlreadyExists {
		return errors.New("password already exists")
	}

	p, err := utils.PrettifyJSON(resp)
	if err != nil {
		return err
	}

	fmt.Println(p)

	return nil
}
