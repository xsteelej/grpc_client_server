package rest

import (
	"context"
	"encoding/json"
	mux "github.com/gorilla/mux"
	grpc "github.com/xsteelej/grpc_client_server/grpc"
	"log"
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
	ctx := context.Background()
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		port, err := s.db.Read(ctx, &grpc.PortRequest{Id: id})
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			_, err := w.Write([]byte(id + " Not found!"))
			if err != nil {
				log.Println("Error writing to http.ResponseWriter")
			}
			return
		}

		portBytes, err := json.Marshal(port)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, err := w.Write([]byte("Internal server error"))
			if err != nil {
				log.Println("Error writing to http.ResponseWriter " + err.Error())
			}
			return
		}

		_, err = w.Write(portBytes)
		if err != nil {
			log.Println("Error writing to http.ResponseWriter " + err.Error())
		}
	}
}
