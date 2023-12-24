package setup

import (
	"context"
	"fmt"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/urfave/cli"
)

func SetupVerifier(ctx context.Context, c *cli.Context) (*oidc.IDTokenVerifier, error) {
	issuer := fmt.Sprintf("%v/dex", c.String("dex-base-url"))

	provider, err := oidc.NewProvider(ctx, issuer)
	if err != nil {
		return nil, err
	}

	idTokenVerifier := provider.Verifier(&oidc.Config{ClientID: c.String("client-id")})

	return idTokenVerifier, nil
}
