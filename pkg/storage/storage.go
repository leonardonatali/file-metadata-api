package storage

import (
	"io"
	"time"
)

type StorageService interface {
	Load(cfg *StorageConfig) error
	BucketExists() (bool, error)
	CreateBucket() error
	PutFile(content io.Reader, path, mimeType string, size uint64) error
	DeleteFile(path string) error
	GetDownloadURL(path, filename, mimeType string, expires time.Duration) (string, error)
}
