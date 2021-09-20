package files

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leonardonatali/file-metadata-api/pkg/auth"
	"github.com/leonardonatali/file-metadata-api/pkg/config"
	"github.com/leonardonatali/file-metadata-api/pkg/files/dto"
	"github.com/leonardonatali/file-metadata-api/pkg/files/repository/postgres"
	"github.com/leonardonatali/file-metadata-api/pkg/users/entities"
	"gorm.io/gorm"
)

type FilesController struct {
	cfg          *config.Config
	filesService *FilesService
}

func NewFilesController(cfg *config.Config, db *gorm.DB) *FilesController {
	return &FilesController{
		cfg:          &config.Config{},
		filesService: NewFilesService(postgres.NewPostgresFilesRepository(db)),
	}
}

func (c *FilesController) UploadFile(ctx *gin.Context) {

	var createFileDto dto.CreateFileDto

	//Validate input
	_, err := ctx.MultipartForm()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": "the request input must be multipart/form"})
		return
	}

	if err := ctx.ShouldBind(&createFileDto); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Extract metadata from file
	createFileDto.LoadMetadata()

	user := ctx.Request.Context().Value(auth.ContextUserKey).(*entities.User)
	createFileDto.UserID = user.ID

	result, err := c.filesService.CreateFile(&createFileDto)
	if err != nil {
		log.Printf("cannot store the file: %s", err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "cannot save the file"})
		return
	}

	ctx.PureJSON(http.StatusCreated, result)
}
