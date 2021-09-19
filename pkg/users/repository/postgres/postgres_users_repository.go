package postgres

import (
	"errors"

	"github.com/leonardonatali/file-metadata-api/pkg/users/entities"
	"gorm.io/gorm"
)

type PostgresUsersRepository struct {
	db *gorm.DB
}

func NewPostgresUsersRepository(db *gorm.DB) *PostgresUsersRepository {
	return &PostgresUsersRepository{
		db: db,
	}
}

func (r *PostgresUsersRepository) CreateUser(user *entities.User) error {
	return r.db.Create(user).Error
}

func (r *PostgresUsersRepository) GetUser(id uint64) (*entities.User, error) {
	var user entities.User

	query := r.db.Where("id = ?", id).First(&user)
	if query.Error != nil {
		if errors.Is(query.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, query.Error
	}

	return &user, nil
}
