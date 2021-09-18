package repository

import "github.com/leonardonatali/file-metadata-api/pkg/files/entities"

type FilesRepository interface {
	CreateFile(file *entities.File) error
	GetAllFiles(userID uint64, path string) ([]*entities.File, error)
	GetFile(fileID uint64) (*entities.File, error)
	GetFileMetadata(fileID uint64) ([]*entities.FileMetadata, error)
	ReplaceFile(oldFile, newFile *entities.File) error
}
