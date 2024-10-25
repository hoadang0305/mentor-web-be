package repository

import (
	"context"
	"github.com/hoadang0305/mentor-web-be/internal/domain/entity"
)

type UserRepository interface {
	GetAllUser(c context.Context) []entity.User
}
