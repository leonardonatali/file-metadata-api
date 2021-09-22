package files

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/leonardonatali/file-metadata-api/pkg/auth"
	"github.com/leonardonatali/file-metadata-api/pkg/config"
	"github.com/leonardonatali/file-metadata-api/pkg/files/dto"
	"github.com/leonardonatali/file-metadata-api/pkg/files/repository/postgres"
	"github.com/leonardonatali/file-metadata-api/pkg/storage"
	"github.com/leonardonatali/file-metadata-api/pkg/users/entities"
	"gorm.io/gorm"
)

type FilesController struct {
	cfg            *config.Config
	filesService   *FilesService
	storageService storage.StorageService
}

func NewFilesController(cfg *config.Config, db *gorm.DB, storageService storage.StorageService) *FilesController {
	return &FilesController{
		cfg:            &config.Config{},
		filesService:   NewFilesService(postgres.NewPostgresFilesRepository(db)),
		storageService: storageService,
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
	createFileDto.Name = dto.GetFilename(createFileDto.File)

	result, err := c.filesService.CreateFile(&createFileDto)
	if err != nil {
		log.Printf("cannot save the file: %s", err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "cannot save the file"})
		return
	}

	//Put file on storage
	file, err := createFileDto.File.Open()
	if err != nil {
		log.Printf("cannot open the file: %s", err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "cannot open the file"})
		return
	}

	err = c.storageService.PutFile(
		file,
		result.GetQualifiedName(),
		result.GetMetadataByKey("type").Value,
		uint64(createFileDto.File.Size),
	)

	if err != nil {
		log.Printf("cannot store the file: %s", err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "cannot store the file"})
		return
	}

	ctx.JSON(http.StatusCreated, result)
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

	url, err := c.storageService.GetDownloadURL(
		file.GetQualifiedName(),
		file.Name,
		file.GetMetadataByKey("type").Value,
		time.Hour,
	)

	if err != nil || file == nil {
		log.Printf("cannot get download URL: %s", err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "cannot get download URL"})
		return
	}

	ctx.PureJSON(http.StatusOK, dto.DownloadFileDtoResponse{DownloadURL: url})
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
	if err := c.storageService.DeleteFile(file.GetQualifiedName()); err != nil {
		log.Printf("cannot delete file from storage: %s", err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "cannot delete file from storage"})
		return
	}

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
	src := file.GetQualifiedName()

	updateFilePathDto.FileID = file.ID

	if err := c.filesService.UpdateFilePath(updateFilePathDto); err != nil {
		log.Printf("cannot change file path: %s", err.Error())
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "cannot change file path"})
		return
	}

	file.Path = updateFilePathDto.Path
	dest := file.GetQualifiedName()

	if err = c.storageService.Move(src, dest); err != nil {
		log.Printf("cannot move file in storage: %s", err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "cannot move file in storage"})
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
		return
	}

	if err := c.storageService.DeleteFile(file.GetQualifiedName()); err != nil {
		log.Printf("cannot delete old file: %s", err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "cannot delete old file"})
		return
	}

	content, err := updateFileDto.File.Open()
	if err != nil {
		log.Printf("cannot read file contents: %s", err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "cannot read file contents"})
		return
	}

	err = c.storageService.PutFile(
		content,
		file.GetQualifiedName(),
		file.GetMetadataByKey("type").Value,
		uint64(updateFileDto.File.Size),
	)

	if err != nil {
		log.Printf("cannot store file in storage: %s", err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "cannot store file in storage"})
		return
	}
}

func (c *FilesController) GetFileTree(ctx *gin.Context) {
	user := ctx.Request.Context().Value(auth.ContextUserKey).(*entities.User)

	paths, _ := c.filesService.GetAllPaths(user.ID)
	paths = getUniques(paths)

	response := []dto.GetFileTreeDto{}
	for i := range paths {
		response = addToTree(response, strings.Split(paths[i], "/"))
	}
	ctx.PureJSON(200, response)
}

func getUniques(items []string) []string {
	uniques := []string{}

	for _, item := range items {
		exists := false
		for _, unique := range uniques {
			if item == unique {
				exists = true
				break
			}
		}
		if !exists {
			uniques = append(uniques, item)
		}
	}

	return uniques
}

func addToTree(root []dto.GetFileTreeDto, names []string) []dto.GetFileTreeDto {
	if len(names) > 0 {
		var i int

		//Verifica se  já não existe um nó filho com o mesmo nome
		for i = 0; i < len(root); i++ {
			if root[i].CurrentDir == names[0] {
				break
			}
		}

		// Caso não exista, adiciona a lista do nó raiz, o novo nó
		if i == len(root) {
			root = append(root, dto.GetFileTreeDto{CurrentDir: names[0]})
		}

		// Chama recursivamente sempre avançando o início da lista
		root[i].Children = addToTree(root[i].Children, names[1:])
	}
	return root
}
