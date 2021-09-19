package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leonardonatali/file-metadata-api/pkg/auth/middlewares"
	"github.com/leonardonatali/file-metadata-api/pkg/config"
	"github.com/leonardonatali/file-metadata-api/pkg/database/migrations"
	"github.com/leonardonatali/file-metadata-api/pkg/files"
	"github.com/leonardonatali/file-metadata-api/pkg/users"
	"github.com/leonardonatali/file-metadata-api/pkg/users/repository/postgres"
	"gorm.io/gorm"
)

type Server struct {
	cfg    *config.Config
	router *gin.Engine
	db     *gorm.DB
}

func NewServer(cfg *config.Config) *Server {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	//Define o tamanho máximo aceito pelo form em 10MB
	router.MaxMultipartMemory = 10 * 1024 * 1024

	return &Server{
		cfg:    cfg,
		router: router,
		db:     cfg.GetDBConn(),
	}
}

func (s *Server) Run() {
	s.Migrate()
	s.RegisterRoutes()
	s.listen()
}

func (s *Server) Migrate() {
	migrations.Migrate(s.cfg)
}

func (s *Server) RegisterRoutes() {
	usersService := users.NewUsersService(postgres.NewPostgresUsersRepository(s.db))
	filesController := files.NewFilesController(s.cfg, s.db)

	root := s.router.Group("/")
	root.Use(middlewares.GetAuthMiddleware(usersService))

	root.GET("/ok", func(c *gin.Context) {
		c.String(http.StatusOK, ":)")
	})

	filesGroup := root.Group("/files")
	filesGroup.POST("/upload", filesController.UploadFile)
}

func (s *Server) listen() {
	addr := fmt.Sprintf(":%d", s.cfg.Port)

	log.Printf("The server is running at %s", addr)
	s.router.Run(addr)
}
