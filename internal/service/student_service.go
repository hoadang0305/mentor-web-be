package service

import (
	"context"
	"github.com/hoadang0305/mentor-web-be/internal/domain/model"
)

type StudentService interface {
	GetAllStudent(ctx context.Context) []model.Student
}
