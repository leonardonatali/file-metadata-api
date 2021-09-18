package repository

import "github.com/leonardonatali/file-metadata-api/pkg/users/entities"

type FilesRepository interface {
	CreateUser(user entities.User) error
	GetUser(id uint64) (*entities.User, error)
	Exists(token string) (bool, error)
}
