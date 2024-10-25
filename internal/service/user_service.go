package service

import (
	"context"
	"github.com/hoadang0305/mentor-web-be/internal/domain/entity"
)

type UserService interface {
	GetAllUser(ctx context.Context) []entity.User
}
