package files

import (
	"log"
	"net/http"
	"strconv"

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

func (c *FilesController) GetAllFiles(ctx *gin.Context) {
	user := ctx.Request.Context().Value(auth.ContextUserKey).(*entities.User)
	path := ctx.Query("path")

	request := &dto.GetFilesDto{UserID: user.ID}
	if path != "" {
		request.Path = path
	}

	files, err := c.filesService.GetAllFiles(request)
	if err != nil {
		log.Printf("cannot get files: %s", err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "cannot get files"})
		return
	}

	ctx.JSON(http.StatusOK, files)
}

func (c *FilesController) GetFileMetadata(ctx *gin.Context) {
	var req dto.GetMetadataDto
	var err error

	req.FileID, err = strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := ctx.Request.Context().Value(auth.ContextUserKey).(*entities.User)
	file, err := c.filesService.GetFile(req.FileID, user.ID)
	if err != nil || file == nil {
		if err != nil {
			log.Printf("cannot get file metadata: %s", err.Error())
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "file not found"})
		return
	}

	ctx.PureJSON(http.StatusOK, file.Metadata)
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
	createFileDto.Metadata = dto.GetMetadata(createFileDto.File, createFileDto.Path)

	user := ctx.Request.Context().Value(auth.ContextUserKey).(*entities.User)
	createFileDto.UserID = user.ID

	result, err := c.filesService.CreateFile(&createFileDto)
	if err != nil {
		log.Printf("cannot store the file: %s", err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "cannot save the file"})
		return
	}

	//Put file in storage
	ctx.PureJSON(http.StatusCreated, result)
}

func (c *FilesController) DownloadFile(ctx *gin.Context) {
	fileId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	user := ctx.Request.Context().Value(auth.ContextUserKey).(*entities.User)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "file id is required"})
		return
	}

	file, err := c.filesService.GetFile(fileId, user.ID)
	if err != nil || file == nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "file not found"})
		return
	}

	// Get download URL from storage
	ctx.SecureJSON(http.StatusOK, dto.DownloadFileDtoResponse{DownloadURL: file.Path})
}

func (c *FilesController) DeleteFile(ctx *gin.Context) {
	user := ctx.Request.Context().Value(auth.ContextUserKey).(*entities.User)
	fileId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "file id is required"})
		return
	}

	file, err := c.filesService.GetFile(fileId, user.ID)
	if err != nil || file == nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "file not found"})
		return
	}

	if err := c.filesService.DeleteFile(&dto.DeleteFileDto{
		UserID: user.ID,
		FileID: file.ID,
	}); err != nil {
		log.Printf("error while deleting file: %s", err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "cannot delete file"})
		return
	}

	// Delete file from storage

	ctx.Status(http.StatusOK)
}

func (c *FilesController) UpdatePath(ctx *gin.Context) {
	fileId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "file id is required"})
		return
	}

	updateFilePathDto := &dto.UpdateFilePathDto{}
	if err := ctx.ShouldBind(updateFilePathDto); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := ctx.Request.Context().Value(auth.ContextUserKey).(*entities.User)

	file, err := c.filesService.GetFile(fileId, user.ID)
	if err != nil || file == nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "file not found"})
		return
	}

	updateFilePathDto.FileID = file.ID

	if err := c.filesService.UpdateFilePath(updateFilePathDto); err != nil {
		log.Printf("cannot change file path: %s", err.Error())
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "cannot change file path"})
		return
	}
}

func (c *FilesController) UpdateFile(ctx *gin.Context) {
	//Validate input
	_, err := ctx.MultipartForm()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": "the request input must be multipart/form"})
		return
	}

	updateFileDto := &dto.UpdateFileDto{}

	if err := ctx.ShouldBind(updateFileDto); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fileId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "file id is required"})
		return
	}

	user := ctx.Request.Context().Value(auth.ContextUserKey).(*entities.User)

	file, err := c.filesService.GetFile(fileId, user.ID)
	if err != nil || file == nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "file not found"})
		return
	}

	//Extract metadata from file
	updateFileDto.Metadata = dto.GetMetadata(updateFileDto.File, file.Path)
	updateFileDto.OldFileID = file.ID
	updateFileDto.UserID = user.ID

	if err := c.filesService.UpdateFile(updateFileDto); err != nil {
		log.Printf("cannot replace file: %s", err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "cannot replace file"})
	}
}
