package files

import (
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

func (s *FilesService) CreateFile(dto *dto.CreateFileDto) (*entities.File, error) {

	file := &entities.File{
		ID:        0,
		UserID:    dto.UserID,
		Name:      dto.Name,
		Path:      dto.Path,
		Metadata:  parseMetadata(dto.Metadata),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := s.filesRepository.CreateFile(file)
	return file, err
}

func (s *FilesService) GetFile(id, userID uint64) (*entities.File, error) {
	return s.filesRepository.GetFile(id, userID)
}

func (s *FilesService) GetAllFiles(dto *dto.GetFilesDto) ([]*entities.File, error) {
	return s.filesRepository.GetAllFiles(dto.UserID, dto.Path)
}

func (s *FilesService) GetFileMetadata(dto *dto.GetMetadataDto) ([]*entities.FilesMetadata, error) {
	return s.filesRepository.GetFileMetadata(dto.FileID)
}

func (s *FilesService) UpdateFile(dto *dto.UpdateFileDto) error {
	return s.filesRepository.UpdateFile(dto.OldFileID, dto.Path, parseMetadata(dto.Metadata))
}

func (s *FilesService) UpdateFilePath(dto *dto.UpdateFilePathDto) error {
	return s.filesRepository.UpdateFilePath(dto.FileID, dto.Path)
}

func (s *FilesService) DeleteFile(dto *dto.DeleteFileDto) error {
	return s.filesRepository.DeleteFile(dto.UserID, dto.FileID)
}

func parseMetadata(content map[string]string) []*entities.FilesMetadata {
	metadata := []*entities.FilesMetadata{}

	for key, value := range content {
		metadata = append(metadata, &entities.FilesMetadata{
			Key:   key,
			Value: value,
		})
	}

	return metadata
}
