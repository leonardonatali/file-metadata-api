package repository

import "github.com/leonardonatali/file-metadata-api/pkg/files/entities"

type FilesRepository interface {
	CreateFile(file *entities.File) error
	GetAllFiles(userID uint64, path string) ([]*entities.File, error)
	GetFile(fileID, userID uint64) (*entities.File, error)
	GetFileMetadata(fileID uint64) ([]*entities.FilesMetadata, error)
	UpdateFile(id uint64, path string, metadata []*entities.FilesMetadata) error
	UpdateFilePath(fileID uint64, path string) error
	DeleteFile(userID, fileID uint64) error
}
