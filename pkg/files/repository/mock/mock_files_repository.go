package mock

import (
	"fmt"

	"github.com/leonardonatali/file-metadata-api/pkg/files/entities"
)

type MockFilesRepository struct {
	EnableErrors     bool
	ReturnEmptyValue bool
}

func getMetadata(fileID uint64, path string) []*entities.FilesMetadata {
	return []*entities.FilesMetadata{
		{ID: 1, FileID: 1, Key: "filename", Value: "file_1.js"},
		{ID: 2, FileID: 1, Key: "path", Value: path},
		{ID: 3, FileID: 1, Key: "size", Value: "5501"},
		{ID: 4, FileID: 1, Key: "type", Value: "application/javascript"},
	}
}

func (r *MockFilesRepository) GetErr(message string) error {
	return fmt.Errorf("[MOCK] MockFilesRepository: %s", message)
}

func (r *MockFilesRepository) CreateFile(file *entities.File) error {
	if r.EnableErrors {
		return r.GetErr("CreateFile")
	}
	return nil
}

func (r *MockFilesRepository) GetAllFiles(userID uint64, path string) ([]*entities.File, error) {
	if r.EnableErrors {
		return nil, r.GetErr("GetAllFiles")
	}

	if r.ReturnEmptyValue {
		return []*entities.File{}, nil
	}

	files := []*entities.File{
		{
			ID:       1,
			UserID:   userID,
			Name:     "name_file.exe",
			Path:     path,
			Metadata: getMetadata(1, path),
		},
		{
			ID:       2,
			UserID:   userID,
			Name:     "name_file.exe(2)",
			Path:     path,
			Metadata: getMetadata(2, path),
		},
	}

	return files, nil
}

func (r *MockFilesRepository) GetFile(fileID, userID uint64) (*entities.File, error) {
	if r.EnableErrors {
		return nil, r.GetErr("GetFile")
	}

	if r.ReturnEmptyValue {
		return nil, nil
	}

	return &entities.File{
		ID:       fileID,
		UserID:   userID,
		Name:     "name_file.exe",
		Path:     "test/path",
		Metadata: getMetadata(fileID, "test/path"),
	}, nil
}

func (r *MockFilesRepository) GetFileMetadata(fileID uint64) ([]*entities.FilesMetadata, error) {
	if r.EnableErrors {
		return nil, r.GetErr("GetFileMetadata")
	}

	if r.ReturnEmptyValue {
		return []*entities.FilesMetadata{}, nil
	}

	return getMetadata(fileID, "test/path"), nil
}

func (r *MockFilesRepository) UpdateFile(id uint64, path string, metadata []*entities.FilesMetadata) error {
	if r.EnableErrors {
		return r.GetErr("UpdateFile")
	}
	return nil
}

func (r *MockFilesRepository) UpdateFilePath(fileID uint64, path string) error {
	if r.EnableErrors {
		return r.GetErr("UpdateFilePath")
	}

	return nil
}

func (r *MockFilesRepository) DeleteFile(userID, fileID uint64) error {
	if r.EnableErrors {
		return r.GetErr("DeleteFile")
	}

	return nil
}

// Busca todos os metadados de um usuário
func (r *MockFilesRepository) GetAllMetadata(userID uint64) ([]*entities.FilesMetadata, error) {
	if r.EnableErrors {
		return nil, r.GetErr("GetAllMetadata")
	}

	if r.ReturnEmptyValue {
		return []*entities.FilesMetadata{}, nil
	}

	result := []*entities.FilesMetadata{}

	result = append(result, getMetadata(1, "test/path")...)
	result = append(result, getMetadata(2, "test/path")...)

	return result, nil
}
