package repositoryimplement

import (
	"context"
	"github.com/hoadang0305/mentor-web-be/internal/database"
	"github.com/hoadang0305/mentor-web-be/internal/domain/entity"
	"github.com/hoadang0305/mentor-web-be/internal/repository"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepository struct {
	db *mongo.Client
}

func NewUserRepository(db database.Db) repository.UserRepository {
	return &UserRepository{db: db}
}

func (repo UserRepository) GetAllUser(c context.Context) ([]entity.User, error) {
	var users []entity.User
	userCollection := database.GetCollection(repo.db, "users")

	data, err := userCollection.Find(c, bson.M{})
	if err != nil {
		return nil, err
	}
	if err = data.All(c, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (repo UserRepository) AddNewUser(c context.Context, user *entity.User) error {
	userCollection := database.GetCollection(repo.db, "users")
	_, err := userCollection.InsertOne(c, user)
	if err != nil {
		return err
	}
	return nil
}
