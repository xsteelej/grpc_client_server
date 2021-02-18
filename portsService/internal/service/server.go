package service

import (
	context "context"
	grpc "github.com/xsteelej/grpc_client_server/grpc"
	"github.com/xsteelej/grpc_client_server/portsService/internal/database"
	"log"
)

type Server struct {
	Repo database.Repository
	grpc.UnimplementedPortsDatabaseServer
}

func NewServiceService(repo database.Repository) *Server {
	s := &Server{}
	s.Repo = repo
	return s
}

func (s *Server) Write(ctx context.Context, port *grpc.Port) (*grpc.WriteResponse, error) {
	log.Println("Writing port id: " + port.Id)
	success, err := s.Repo.Write(ctx,port)
	if err != nil {
		return nil, err
	}
	return &grpc.WriteResponse{
		Id:      port.Id,
		Success: success,
	}, nil
}

func (s *Server) Read(ctx context.Context, request *grpc.PortRequest) (*grpc.Port, error) {
	log.Println("Reading port id: " + request.Id)
	return s.Repo.Read(ctx, request.Id)
}

