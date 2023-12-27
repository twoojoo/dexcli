package base

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/twoojoo/dexctl/server"
	"github.com/twoojoo/dexctl/setup"
	"github.com/twoojoo/dexctl/utils"
	"github.com/urfave/cli"
)

var SignonFlags []cli.Flag = []cli.Flag{
	cli.StringFlag{
		Name:  "dex-base-url",
		Value: "http://127.0.0.1:5556",
	},
	cli.UintFlag{
		Name:  "port, p",
		Value: 5555,
	},
	cli.StringFlag{
		Name:  "secret, s",
		Value: "example-app-secret",
	},
	cli.StringFlag{
		Name:  "state, t",
		Value: "default-state",
	},
	cli.BoolFlag{
		Name: "userinfo, u",
	},
	cli.BoolFlag{
		Name: "offline-access, o",
	},
	cli.StringSliceFlag{
		Name:  "scope",
		Value: &cli.StringSlice{"profile", "email"},
	},
	cli.StringFlag{
		Name:     "browser, b",
		Required: false,
	},
}

func Signon(c *cli.Context) error {
	ctx := context.Background()

	clientID := c.Args().Get(0)
	if clientID == "" {
		return errors.New("client id must be provided as first argument")
	}

	provider, verifier, config, err := setup.SetupProvider(ctx, clientID, c)
	if err != nil {
		return err
	}

	browser := c.String("browser")
	if browser == "" {
		browser, err = utils.GetDefaultBrowser()
		if err != nil {
			return fmt.Errorf("can't find default browser, consider using the --browser option")
		}
	}

	go func() {
		time.Sleep(200 * time.Microsecond)
		err := utils.RunCommand(browser, fmt.Sprintf("http://localhost:%v/login", c.Uint("port")))
		if err != nil {
			log.Fatal(err)
		}
	}()

	err = server.RunServer(ctx, c, provider, verifier, config, c.String("state"))
	if err != nil {
		return err
	}

	return nil
}
