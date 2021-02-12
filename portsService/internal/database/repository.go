package database

import portsDB "portsService/grpc"
import "context"

type Repository interface {
	Write(ctx context.Context, port *portsDB.Port) (bool, error)
	Read(ctx context.Context, id string) (*portsDB.Port, error)
}