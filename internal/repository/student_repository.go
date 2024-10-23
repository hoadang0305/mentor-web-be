package repository

import (
	"context"
	"github.com/hoadang0305/mentor-web-be/internal/domain/entity"
)

type StudentRepository interface {
	GetAllStudentQuery(c context.Context) []entity.Student
}
