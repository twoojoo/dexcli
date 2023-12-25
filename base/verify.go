package base

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/twoojoo/dexctl/setup"
	"github.com/twoojoo/dexctl/utils"
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
	cli.StringFlag{
		Name:     "field, f",
		Required: false,
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

	clientID := c.Args().Get(0)
	if clientID == "" {
		return errors.New("a token must be provided as second argument")
	}

	verifier, _, err := setup.SetupProvider(ctx, clientID, c)
	if err != nil {
		return err
	}

	rawIDToken := c.Args().Get(1)
	if rawIDToken == "" {
		return errors.New("a token must be provided as second argument")
	}

	idToken, err := verifier.Verify(ctx, rawIDToken)
	if err != nil {
		return err
	}

	claims := Claims{}
	if err := idToken.Claims(&claims); err != nil {
		return err
	}

	field := strings.Title(c.String("field"))
	if field != "" {
		if value, ok := utils.GetStructField(claims, field); ok {
			fmt.Print(value)
			return nil
		}

		return fmt.Errorf("field %v not in result", field)
	}

	p, err := utils.PrettifyJSON(claims)
	if err != nil {
		return err
	}

	fmt.Println(p)

	return nil
}
