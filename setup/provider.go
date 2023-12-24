package setup

import (
	"context"
	"fmt"
	"log"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/urfave/cli"
	"golang.org/x/oauth2"
)

type ProviderParams struct {
	DexURL       string
	ClientID     string
	ClientSecret string
	Port         uint
	Scopes       []string
}

func SetupProvider(ctx context.Context, c *cli.Context) (*oidc.IDTokenVerifier, oauth2.Config, error) {
	values := ProviderParams{
		DexURL:       c.String("dex-url"),
		Port:         c.Uint("port"),
		ClientID:     c.String("client-id"),
		ClientSecret: c.String("client-secret"),
		Scopes:       c.StringSlice("scopes"),
	}

	issuer := fmt.Sprintf("%v/dex", values.DexURL)
	log.Println(issuer)

	provider, err := oidc.NewProvider(ctx, issuer)
	if err != nil {
		return nil, oauth2.Config{}, err
	}

	// Create an ID token parser.
	idTokenVerifier := provider.Verifier(&oidc.Config{ClientID: values.ClientID})

	// Configure the OAuth2 config with the client values.
	oauth2Config := oauth2.Config{
		ClientID:     values.ClientID,
		ClientSecret: values.ClientSecret,
		RedirectURL:  fmt.Sprintf("http://127.0.0.1:%v/callback", values.Port),
		Endpoint:     provider.Endpoint(),
		Scopes:       append([]string{oidc.ScopeOpenID}, values.Scopes...),
	}

	return idTokenVerifier, oauth2Config, nil
}
