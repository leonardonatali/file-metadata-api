package files

import (
	"fmt"
	"time"

	"github.com/leonardonatali/file-metadata-api/pkg/files/dto"
	"github.com/leonardonatali/file-metadata-api/pkg/files/entities"
	"github.com/leonardonatali/file-metadata-api/pkg/files/repository"
)

type FilesService struct {
	filesRepository repository.FilesRepository
}

func NewFilesService(filesRepository repository.FilesRepository) *FilesService {
	return &FilesService{
		filesRepository: filesRepository,
	}
}

func (s *FilesService) CreateFile(dto *dto.CreateFileDto) error {

	file := &entities.File{
		ID:        0,
		UserID:    dto.UserID,
		Name:      dto.Name,
		Path:      dto.Path,
		Metadata:  parseMetadata(dto.Metadata),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return s.filesRepository.CreateFile(file)
}

func (s *FilesService) GetAllFiles(dto *dto.GetFilesDto) ([]*entities.File, error) {
	return s.filesRepository.GetAllFiles(dto.UserID, dto.Path)
}

func (s *FilesService) GetFileMetadata(dto *dto.GetMetadataDto) ([]*entities.FileMetadata, error) {
	return s.filesRepository.GetFileMetadata(dto.FileID)
}

func (s *FilesService) ReplaceFile(dto *dto.ReplaceFileDto) error {

	newFile := &entities.File{
		ID:       0,
		UserID:   dto.UserID,
		Name:     dto.Name,
		Path:     dto.Path,
		Metadata: parseMetadata(dto.Metadata),
	}

	oldFile, err := s.filesRepository.GetFile(dto.OldFileID)
	if err != nil || oldFile == nil {
		return fmt.Errorf("cannot find the old file to be replaced")
	}

	if err := s.filesRepository.ReplaceFile(oldFile, newFile); err != nil {
		return fmt.Errorf("error while replacing file")
	}

	return nil
}

func parseMetadata(content map[string]string) []*entities.FileMetadata {
	metadata := []*entities.FileMetadata{}

	for key, value := range content {
		metadata = append(metadata, &entities.FileMetadata{
			Key:   key,
			Value: value,
		})
	}

	return metadata
}
