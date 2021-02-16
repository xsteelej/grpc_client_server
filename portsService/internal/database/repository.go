package database

import grpc "github.com/xsteelej/grpc_client_server/grpc"
import "context"

type Repository interface {
	Write(ctx context.Context, port *grpc.Port) (bool, error)
	Read(ctx context.Context, id string) (*grpc.Port, error)
}