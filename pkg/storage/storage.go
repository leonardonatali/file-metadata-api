package storage

import "io"

type StorageService interface {
	Load(cfg *StorageConfig) error
	BucketExists() (bool, error)
	CreateBucket(name string) error
	PutFile(content io.Reader, path, mimeType string, size uint64) error
	DeleteFile(path string) error
	GetDownloadURL() (string, error)
}
