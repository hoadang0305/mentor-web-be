package repositoryimplement

import (
	"context"
	"github.com/hoadang0305/mentor-web-be/internal/database"
	"github.com/hoadang0305/mentor-web-be/internal/domain/entity"
	"github.com/hoadang0305/mentor-web-be/internal/repository"
	"github.com/jmoiron/sqlx"
)

type StudentRepository struct {
	db *sqlx.DB
}

func NewStudentRepository(db database.Db) repository.StudentRepository {
	return &StudentRepository{db: db}
}

func (repo StudentRepository) GetAllStudentQuery(c context.Context) []entity.Student {
	return []entity.Student{
		{Name: "John"},
		{Name: "Khoa"},
		{Name: "Lindan"},
	}
}
