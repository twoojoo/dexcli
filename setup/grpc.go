package setup

import (
	"context"

	"fmt"

	"github.com/dexidp/dex/api/v2"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func SetupGrpcClient(ctx context.Context, c *cli.Context) (api.DexClient, error) {
	url := c.String("grpc-url")

	fmt.Println(url)
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("dial: %v", err)
	}

	return api.NewDexClient(conn), nil
}
