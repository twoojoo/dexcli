package token

import (
	"context"
	"errors"
	"fmt"

	"github.com/twoojoo/dexctl/setup"
	"github.com/twoojoo/dexctl/utils"
	"github.com/urfave/cli"
)

var VerifyTokenFlags []cli.Flag = []cli.Flag{
	cli.StringFlag{
		Name:  "dex-base-url",
		Value: "http://127.0.0.1:5556",
	},
}

func VerifyToken(c *cli.Context) error {
	ctx := context.Background()

	clientID := c.Args().Get(0)
	if clientID == "" {
		return errors.New("the client_id must be provided as second argument")
	}

	_, verifier, _, err := setup.SetupProvider(ctx, clientID, c)
	if err != nil {
		return err
	}

	rawToken := c.Args().Get(1)
	if rawToken == "" {
		return errors.New("the token must be provided as second argument")
	}

	token, err := verifier.Verify(ctx, rawToken)
	if err != nil {
		return err
	}

	p, err := utils.PrettifyJSON(token)
	if err != nil {
		return err
	}

	fmt.Println(p)

	return nil
}
