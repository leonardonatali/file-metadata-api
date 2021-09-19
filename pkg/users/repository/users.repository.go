package repository

import "github.com/leonardonatali/file-metadata-api/pkg/users/entities"

type UsersRepository interface {
	CreateUser(user *entities.User) (*entities.User, error)
	GetUser(id uint64, token string) (*entities.User, error)
}
