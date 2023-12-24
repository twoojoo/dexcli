package signin

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

func Signin(c *cli.Context) error {
	ctx := context.Background()

	provider, verifier, err := setup.SetupProvider(ctx, c)
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
		log.Println("launching browser")
		err := utils.RunCommand(browser, fmt.Sprintf("http://localhost:%v/login", c.Uint("port")))
		if err != nil {
			log.Fatal(err)
		}
	}()

	err = server.RunServer(ctx, c, provider, verifier, c.String("state"))
	if err != nil {
		return err
	}

	return nil
}
