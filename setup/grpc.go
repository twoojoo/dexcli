package setup

import (
	"context"

	"fmt"
	"github.com/dexidp/dex/api/v2"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
)

func SetupGrpcClient(ctx context.Context, c *cli.Context) (api.DexClient, error) {
	conn, err := grpc.Dial(c.String("grpc-url"))
	if err != nil {
		return nil, fmt.Errorf("dial: %v", err)
	}

	return api.NewDexClient(conn), nil
}
