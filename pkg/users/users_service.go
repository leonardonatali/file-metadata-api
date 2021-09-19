package files

import (
	"github.com/leonardonatali/file-metadata-api/pkg/users/dto"
	"github.com/leonardonatali/file-metadata-api/pkg/users/entities"
	"github.com/leonardonatali/file-metadata-api/pkg/users/repository"
)

type UsersService struct {
	usersRepository repository.UsersRepository
}

func NewFilesService(usersRepository repository.UsersRepository) *UsersService {
	return &UsersService{
		usersRepository: usersRepository,
	}
}

func (s *UsersService) CreateUser(dto *dto.CreateUserDto) error {
	return s.usersRepository.CreateUser(&entities.User{
		Token: dto.Token,
	})
}

func (s *UsersService) GetUser(dto *dto.GetUserDto) (*entities.User, error) {
	return s.usersRepository.GetUser(dto.ID, dto.Token)
}
