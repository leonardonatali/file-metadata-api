package mock

import (
	"fmt"
	"io"
	"time"

	"github.com/leonardonatali/file-metadata-api/pkg/storage"
)

type MockStorageService struct {
	cfg               *storage.StorageConfig
	ReturnErrors      bool
	ReturnEmptyValues bool
}

func (m *MockStorageService) GetErr(msg string) error {
	return fmt.Errorf("[MOCK] MockStorageRepository: %s", msg)
}

func (m *MockStorageService) Load(cfg *storage.StorageConfig) error {
	if m.ReturnErrors {
		return m.GetErr("Load")
	}

	m.cfg = cfg
	return nil
}

func (m *MockStorageService) BucketExists() (bool, error) {
	if m.ReturnErrors {
		return false, m.GetErr("BucketExists")
	}

	if m.ReturnEmptyValues {
		return false, nil
	}

	return true, nil
}

func (m *MockStorageService) CreateBucket() error {
	if m.ReturnErrors {
		return m.GetErr("CreateBucket")
	}

	return nil
}

func (m *MockStorageService) PutFile(content io.Reader, path, mimeType string, size uint64) error {
	if m.ReturnErrors {
		return m.GetErr("PutFile")
	}

	return nil
}

func (m *MockStorageService) DeleteFile(path string) error {
	if m.ReturnErrors {
		return m.GetErr("DeleteFile")
	}

	return nil
}

func (m *MockStorageService) Move(src, dest string) error {
	if m.ReturnErrors {
		return m.GetErr("Move")
	}

	return nil
}

func (m *MockStorageService) GetDownloadURL(path, filename, mimeType string, expires time.Duration) (string, error) {
	if m.ReturnErrors {
		return "", m.GetErr("GetDownloadURL")
	}

	if m.ReturnEmptyValues {
		return "", nil
	}

	return fmt.Sprintf("%s/%s", path, filename), nil
}
