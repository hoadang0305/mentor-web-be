package serviceimplement

import (
	"context"
	"github.com/hoadang0305/mentor-web-be/internal/domain/entity"
	"github.com/hoadang0305/mentor-web-be/internal/domain/model"
	"github.com/hoadang0305/mentor-web-be/internal/repository"
	"github.com/hoadang0305/mentor-web-be/internal/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) service.UserService {
	return &UserService{userRepository: userRepository}
}

func (service UserService) GetAllUser(ctx context.Context) ([]entity.User, error) {
	return service.userRepository.GetAllUser(ctx)
}

func (service UserService) CreateUser(ctx context.Context, userRequest model.UserRequest) (*entity.User, error) {
	user := &entity.User{
		Id:    primitive.NewObjectID(),
		Name:  userRequest.Name,
		Role:  userRequest.Role,
		Email: userRequest.Email,
	}
	err := service.userRepository.AddNewUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
