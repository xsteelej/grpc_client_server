package rest

import (
	mux "github.com/gorilla/mux"
	grpc "github.com/xsteelej/grpc_client_server/grpc"
	"net/http"
)

type Server struct {
	db     grpc.PortsDatabaseClient
	router *mux.Router
}

func NewServer(client grpc.PortsDatabaseClient) *Server {
	s := &Server{
		db: client,
	}
	s.router = mux.NewRouter().StrictSlash(true)
	s.routes()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) handleGetPort() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Here"))
	}
}
