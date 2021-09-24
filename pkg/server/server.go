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
	"github.com/leonardonatali/file-metadata-api/pkg/storage/s3"
	"github.com/leonardonatali/file-metadata-api/pkg/users"
	"github.com/leonardonatali/file-metadata-api/pkg/users/repository/postgres"
	"gorm.io/gorm"
)

type Server struct {
	Cfg        *config.Config
	Storage    storage.StorageService
	StorageCfg *storage.StorageConfig
	Router     *gin.Engine
	Db         *gorm.DB
}

func NewServer(cfg *config.Config, storageCfg *storage.StorageConfig) *Server {
	//Carrega a configuração do storage
	storage := s3.S3Service{}

	return &Server{
		Cfg:        cfg,
		Storage:    &storage,
		StorageCfg: storageCfg,
		Db:         cfg.GetDBConn(),
	}
}

func (s *Server) Run() {
	s.Setup()
	s.Listen()
}

func (s *Server) Setup() {
	s.Migrate()
	s.SetupRouter()
	s.RegisterRoutes()
	s.SetupStorage()
}

func (s *Server) Migrate() {
	migrations.Migrate(s.Cfg)
}

func (s *Server) SetupRouter() {
	gin.SetMode(gin.ReleaseMode)
	if s.Cfg.Debug {
		s.Router = gin.Default()
	} else {
		s.Router = gin.New()
	}

	//Define o tamanho máximo aceito pelo form em 10MB
	s.Router.MaxMultipartMemory = 10 * 1024 * 1024

}

func (s *Server) RegisterRoutes() {

	usersService := users.NewUsersService(postgres.NewPostgresUsersRepository(s.Db))
	filesController := files.NewFilesController(s.Cfg, s.Db, s.Storage)

	root := s.Router.Group("/")
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

func (s *Server) SetupStorage() {
	if err := s.Storage.Load(s.StorageCfg); err != nil {
		log.Fatalf("cannot load storage config: %s", err.Error())
	}

	exists, err := s.Storage.BucketExists()
	if err != nil {
		log.Fatalf("cannot check if bucket exists: %s", err.Error())
	}

	if !exists {
		if err := s.Storage.CreateBucket(); err != nil {
			log.Fatalf("cannot create bucket: %s", err.Error())
		}
	}
}

func (s *Server) Listen() {
	addr := fmt.Sprintf(":%d", s.Cfg.Port)

	log.Printf("The server is running at %s", addr)
	s.Router.Run(addr)
}
