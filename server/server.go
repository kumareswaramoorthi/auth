package server

import (
	"auth/handler"
	"auth/middleware"
	"net/http"
)

type Server struct {
	Router *http.ServeMux
}

func (s *Server) InitRoute(h *handler.Handler) {
	s.Router.HandleFunc("/auth", middleware.JSONandCORS(h.FindUser))
}
