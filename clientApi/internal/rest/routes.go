package rest

import "net/http"

func (s *Server) routes() {
	s.router.Methods(http.MethodGet).Path("/ports/{id}").HandlerFunc(s.handleGetPort())
}
