package postgres

import (
	"github.com/leonardonatali/file-metadata-api/pkg/files/entities"
	"gorm.io/gorm"
)

type PostgresFilesRepository struct {
	db *gorm.DB
}

func NewPostgresFilesRepository(db *gorm.DB) *PostgresFilesRepository {
	return &PostgresFilesRepository{
		db: db,
	}
}

func (r *PostgresFilesRepository) CreateFile(file *entities.File) error {
	return r.db.Create(file).Error
}

func (r *PostgresFilesRepository) GetAllFiles(userID uint64, path string) ([]*entities.File, error) {
	files := []*entities.File{}
	query := r.db.Where("user_id = ?", userID).Find(&files)
	return files, query.Error
}

func (r *PostgresFilesRepository) GetFileMetadata(fileID uint64) ([]*entities.FileMetadata, error) {
	metadata := []*entities.FileMetadata{}
	query := r.db.Where("file_id = ?", fileID).Find(&metadata)
	return metadata, query.Error
}

func (r *PostgresFilesRepository) ReplaceFile(oldFile, newFile *entities.File) error {
	newFile.ID = 0
	return r.db.Model(oldFile).Updates(newFile).Error
}
