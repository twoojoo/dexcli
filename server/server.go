package server

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/urfave/cli"
	"golang.org/x/oauth2"
)

func RunServer(ctx context.Context, c *cli.Context, idTokenVerifier *oidc.IDTokenVerifier, oauth2Config oauth2.Config, state string) error {
	l, err := net.Listen("tcp", fmt.Sprintf(":%v", c.Uint("port")))
	if err != nil {
		return err
	}

	return http.Serve(l, &ApplicationHanlder{
		idTokenVerifier: idTokenVerifier,
		oauth2Config:    oauth2Config,
	})
}
