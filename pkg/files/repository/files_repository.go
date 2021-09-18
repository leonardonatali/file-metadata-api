package repository

import "github.com/leonardonatali/file-metadata-api/pkg/files/entities"

type FilesRepository interface {
	CreateFile(file *entities.File) error
	GetAllFiles(userID uint64, path string) ([]*entities.File, error)
	GetFileMetadata(id uint64) (interface{}, error)
	ReplaceFile(oldFile, newFile *entities.File) error
}
