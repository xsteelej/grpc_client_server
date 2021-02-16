package ports

import (
	"context"
	dbgrpc "github.com/xsteelej/grpc_client_server/grpc"
	"log"
)

func ReadJsonFile(ctx context.Context, filepath string, dbClient dbgrpc.PortsDatabaseClient) {
	cancelled := false
	for !cancelled {
		select {
		case <-ctx.Done():
			log.Println("cancelling..")
			cancelled = true
			break
		default:
			log.Println("readJsonfile")
		}
	}

}
