package base

import (
	"context"
	"errors"
	"fmt"

	"github.com/twoojoo/dexcli/setup"
	"github.com/twoojoo/dexcli/utils"
	"github.com/urfave/cli"
)

var VerifyFlags []cli.Flag = []cli.Flag{
	cli.StringFlag{
		Name:  "dex-base-url",
		Value: "http://127.0.0.1:5556",
	},
	cli.StringFlag{
		Name:  "client-id, cid",
		Value: "example-app",
	},
}

type Claims struct {
	AtHash        string `json:"at_hash"`
	Aud           string `json:"aud"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Iat           int    `json:"iat"`
	Iss           string `json:"iss"`
	Name          string `json:"name"`
	Sub           string `json:"sub"`
}

func Verify(c *cli.Context) error {
	ctx := context.Background()

	verifier, _, err := setup.SetupProvider(ctx, c)
	if err != nil {
		return err
	}

	rawIDToken := c.Args().First()
	if rawIDToken == "" {
		return errors.New("a token must be provided as argument")
	}

	idToken, err := verifier.Verify(ctx, rawIDToken)
	if err != nil {
		return err
	}

	claims := Claims{}
	if err := idToken.Claims(&claims); err != nil {
		return err
	}

	p, err := utils.PrettifyJSON(claims)
	if err != nil {
		return err
	}

	fmt.Println(p)

	return nil
}
