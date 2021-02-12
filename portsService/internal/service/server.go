package service

import (
	context "context"
	portsDB "github.com/xsteelej/grpc_client_server/grpc"
	"github.com/xsteelej/grpc_client_server/portsService/internal/database"
)

type Server struct {
	repo database.Repository
	portsDB.UnimplementedPortsDatabaseServer
}

func (s Server) Write(ctx context.Context, port *portsDB.Port) (*portsDB.WriteResponse, error) {
	success, err := s.repo.Write(ctx,port)
	if err != nil {
		return nil, err
	}
	return &portsDB.WriteResponse{
		Id:      port.Id,
		Success: success,
	}, nil
}

func (s Server) Read(ctx context.Context, request *portsDB.PortRequest) (*portsDB.Port, error) {
	return s.repo.Read(ctx, request.Id)
}

