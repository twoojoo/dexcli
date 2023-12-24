package base

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/twoojoo/dexcli/server"
	"github.com/twoojoo/dexcli/setup"
	"github.com/twoojoo/dexcli/utils"
	"github.com/urfave/cli"
)

var SigninFlags []cli.Flag = []cli.Flag{
	cli.StringFlag{
		Name:  "dex-base-url",
		Value: "http://127.0.0.1:5556",
	},
	cli.UintFlag{
		Name:  "port, p",
		Value: 3000,
	},
	cli.StringFlag{
		Name:  "client-id, cid",
		Value: "example-app",
	},
	cli.StringFlag{
		Name:  "client-secret, csec",
		Value: "example-app-secret",
	},
	cli.StringFlag{
		Name:  "state, t",
		Value: "",
	},
	cli.StringSliceFlag{
		Name:  "scopes, s",
		Value: &cli.StringSlice{"profile", "email"},
	},
	cli.StringFlag{
		Name:     "browser, b",
		Required: false,
	},
}

func Signin(c *cli.Context) error {
	ctx := context.Background()

	verifier, config, err := setup.SetupProvider(ctx, c)
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

	err = server.RunServer(ctx, c, verifier, config, c.String("state"))
	if err != nil {
		return err
	}

	return nil
}
