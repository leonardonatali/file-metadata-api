package server

import (
	"github.com/leonardonatali/file-metadata-api/pkg/config"
	"github.com/leonardonatali/file-metadata-api/pkg/database/migrations"
)

type Server struct {
	cfg *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) Run() {
	migrations.Migrate(s.cfg)

	//Create controllers
	//Register Routes
	//Run server
}
