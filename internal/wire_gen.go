// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package internal

import (
	"github.com/hoadang0305/mentor-web-be/internal/controller"
	"github.com/hoadang0305/mentor-web-be/internal/controller/http"
	"github.com/hoadang0305/mentor-web-be/internal/controller/http/v1"
	"github.com/hoadang0305/mentor-web-be/internal/database"
	"github.com/hoadang0305/mentor-web-be/internal/repository/implement"
	"github.com/hoadang0305/mentor-web-be/internal/service/implement"
	"github.com/google/wire"
)

// Injectors from wire.go:

func InitializeContainer(db database.Db) *controller.ApiContainer {
	studentRepository := repositoryimplement.NewStudentRepository(db)
	studentService := serviceimplement.NewStudentService(studentRepository)
	studentHandler := v1.NewStudentHandler(studentService)
	server := http.NewServer(studentHandler)
	apiContainer := controller.NewApiContainer(server)
	return apiContainer
}

// wire.go:

var container = wire.NewSet(controller.NewApiContainer)

// may have grpc server in the future
var serverSet = wire.NewSet(http.NewServer)

// handler === controller | with service and repository layers to form 3 layers architecture
var handlerSet = wire.NewSet(v1.NewStudentHandler)

var serviceSet = wire.NewSet(serviceimplement.NewStudentService)

var repositorySet = wire.NewSet(repositoryimplement.NewStudentRepository)
