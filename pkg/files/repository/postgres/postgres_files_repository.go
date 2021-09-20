package postgres

import (
	"fmt"
	"strings"

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

	path = strings.TrimPrefix(path, "/")
	path = strings.TrimSuffix(path, "/")

	query := r.db.Where("user_id = ?", userID)

	if path != "" {
		path = fmt.Sprintf("/%s/", path)
		query = query.Where("path LIKE('?%')", path)
	}

	query.Preload("Metadata").Find(&files)
	return files, query.Error
}

func (r *PostgresFilesRepository) GetFile(fileID, userID uint64) (*entities.File, error) {
	file := entities.File{}
	query := r.db

	if fileID > 0 {
		query = query.Where("id = ?", fileID)
	}

	if userID > 0 {
		query = query.Where("user_id = ?", userID)
	}

	if err := query.Preload("Metadata").First(&file).Error; err != nil {
		return nil, err
	}

	return &file, nil
}

func (r *PostgresFilesRepository) GetFileMetadata(fileID uint64) ([]*entities.FilesMetadata, error) {
	metadata := []*entities.FilesMetadata{}
	query := r.db.Where("file_id = ?", fileID).Find(&metadata)
	return metadata, query.Error
}

func (r *PostgresFilesRepository) UpdateFile(id uint64, path string, metadata []*entities.FilesMetadata) error {
	for _, m := range metadata {
		m.FileID = id
	}

	model := &entities.File{ID: id}

	if err := r.db.
		Where("file_id = ?", id).
		Delete(&entities.FilesMetadata{}).
		Error; err != nil {
		return err
	}

	if err := r.db.
		Create(metadata).
		Error; err != nil {
		return err
	}

	if err := r.db.
		Model(model).
		Updates(&entities.File{Path: path}).
		Error; err != nil {
		return err
	}

	return nil
}

func (r *PostgresFilesRepository) UpdateFilePath(fileID uint64, path string) error {
	return r.db.
		Model(&entities.File{ID: fileID}).
		Update("path", path).
		Error
}

func (r *PostgresFilesRepository) DeleteFile(userID, fileID uint64) error {
	return r.db.Model(&entities.File{}).Delete(&entities.File{
		ID:     fileID,
		UserID: userID,
	}).Error
}
