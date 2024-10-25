package serviceimplement

import (
	"context"
	"github.com/hoadang0305/mentor-web-be/internal/domain/entity"
	"github.com/hoadang0305/mentor-web-be/internal/repository"
	"github.com/hoadang0305/mentor-web-be/internal/service"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) service.UserService {
	return &UserService{userRepository: userRepository}
}

func (service UserService) GetAllUser(ctx context.Context) []entity.User {
	usersFromRepo := service.userRepository.GetAllUser(ctx)

	users := make([]entity.User, len(usersFromRepo))
	for i, s := range usersFromRepo {
		users[i] = entity.User{
			Name: s.Name,
		}
	}
	return users
}
