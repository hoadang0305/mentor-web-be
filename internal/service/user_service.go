package service

import (
	"context"
	"github.com/hoadang0305/mentor-web-be/internal/domain/entity"
	"github.com/hoadang0305/mentor-web-be/internal/domain/model"
)

type UserService interface {
	GetAllUser(ctx context.Context) ([]entity.User, error)
	CreateUser(ctx context.Context, userRequest model.UserRequest) (*entity.User, error)
}
