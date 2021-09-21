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
	"github.com/leonardonatali/file-metadata-api/pkg/storage"
	minio "github.com/leonardonatali/file-metadata-api/pkg/storage/storage_minio"
	"github.com/leonardonatali/file-metadata-api/pkg/users"
	"github.com/leonardonatali/file-metadata-api/pkg/users/repository/postgres"
	"gorm.io/gorm"
)

type Server struct {
	cfg        *config.Config
	storage    storage.StorageService
	storageCfg *storage.StorageConfig
	router     *gin.Engine
	db         *gorm.DB
}

func NewServer(cfg *config.Config, storageCfg *storage.StorageConfig) *Server {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	//Define o tamanho máximo aceito pelo form em 10MB
	router.MaxMultipartMemory = 10 * 1024 * 1024

	//Carrega a configuração do storage
	minioService := minio.MinioService{}

	return &Server{
		cfg:        cfg,
		storage:    &minioService,
		storageCfg: storageCfg,
		router:     router,
		db:         cfg.GetDBConn(),
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

	if err := s.storage.Load(s.storageCfg); err != nil {
		log.Fatalf("cannot load storage config: %s", err.Error())
	}

	exists, err := s.storage.BucketExists()
	if err != nil {
		log.Fatalf("cannot check if storage exists: %s", err.Error())
	}

	if !exists {
		if err := s.storage.CreateBucket(); err != nil {
			log.Fatalf("cannot check if storage exists: %s", err.Error())
		}
	}

	usersService := users.NewUsersService(postgres.NewPostgresUsersRepository(s.db))
	filesController := files.NewFilesController(s.cfg, s.db, s.storage)

	root := s.router.Group("/")
	root.Use(middlewares.GetAuthMiddleware(usersService))

	root.GET("/ok", func(c *gin.Context) {
		c.String(http.StatusOK, ":)")
	})

	files := root.Group("/files")
	files.GET("/filetree", filesController.GetFileTree)
	files.POST("/upload", filesController.UploadFile)
	files.GET("/:id/download", filesController.DownloadFile)
	files.GET("/:id/metadata", filesController.GetFileMetadata)
	files.DELETE("/:id", filesController.DeleteFile)
	files.PATCH("/:id", filesController.UpdatePath)
	files.PUT("/:id", filesController.UpdateFile)
	files.GET("/", filesController.GetAllFiles)
}

func (s *Server) listen() {
	addr := fmt.Sprintf(":%d", s.cfg.Port)

	log.Printf("The server is running at %s", addr)
	s.router.Run(addr)
}
