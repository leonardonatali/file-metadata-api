package server

import "github.com/leonardonatali/file-metadata-api/pkg/config"

type Server struct {
	cfg *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) Run() {
	//Create controllers
	//Register Routes
	//Run server
}
