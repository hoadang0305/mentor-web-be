package repositoryimplement

import (
	"context"
	"github.com/hoadang0305/mentor-web-be/internal/database"
	"github.com/hoadang0305/mentor-web-be/internal/domain/entity"
	"github.com/hoadang0305/mentor-web-be/internal/repository"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepository struct {
	db *mongo.Client
}

func NewUserRepository(db database.Db) repository.UserRepository {
	return &UserRepository{db: db}
}

func (repo UserRepository) GetAllUser(c context.Context) []entity.User {
	return []entity.User{
		{Name: "John"},
		{Name: "Khoa"},
	}
}
