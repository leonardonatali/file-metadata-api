package postgres

import (
	"errors"
	"fmt"

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

func (r *PostgresUsersRepository) CreateUser(user *entities.User) (*entities.User, error) {
	dest := &entities.User{}
	query := r.db.FirstOrCreate(dest, &entities.User{Token: user.Token})
	return dest, query.Error
}

func (r *PostgresUsersRepository) GetUser(id uint64, token string) (*entities.User, error) {
	if id <= 0 && token == "" {
		return nil, fmt.Errorf("id or token must be present")
	}

	var user entities.User

	query := r.db

	if token != "" {
		query = query.Where("token = ?", token)
	} else {
		if id > 0 {
			query = query.Where("id = ?", id)
		}
	}

	query = query.First(&user)

	if query.Error != nil {
		if errors.Is(query.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, query.Error
	}

	return &user, nil
}
