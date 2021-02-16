package service

import (
	context "context"
	grpc "github.com/xsteelej/grpc_client_server/grpc"
	"github.com/xsteelej/grpc_client_server/portsService/internal/database"
)

type Server struct {
	repo database.Repository
	grpc.UnimplementedPortsDatabaseServer
}

func (s Server) Write(ctx context.Context, port *grpc.Port) (*grpc.WriteResponse, error) {
	success, err := s.repo.Write(ctx,port)
	if err != nil {
		return nil, err
	}
	return &grpc.WriteResponse{
		Id:      port.Id,
		Success: success,
	}, nil
}

func (s Server) Read(ctx context.Context, request *grpc.PortRequest) (*grpc.Port, error) {
	return s.repo.Read(ctx, request.Id)
}

