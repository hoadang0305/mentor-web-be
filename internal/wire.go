package internal

import (
	"github.com/google/wire"
	"github.com/hoadang0305/mentor-web-be/internal/controller"
	"github.com/hoadang0305/mentor-web-be/internal/controller/http"
	v1 "github.com/hoadang0305/mentor-web-be/internal/controller/http/v1"
	"github.com/hoadang0305/mentor-web-be/internal/database"
	repositoryimplement "github.com/hoadang0305/mentor-web-be/internal/repository/implement"
	serviceimplement "github.com/hoadang0305/mentor-web-be/internal/service/implement"
)

var container = wire.NewSet(
	controller.NewApiContainer,
)

// may have grpc server in the future
var serverSet = wire.NewSet(
	http.NewServer,
)

// handler === controller | with service and repository layers to form 3 layers architecture
var handlerSet = wire.NewSet(
	v1.NewUserHandler,
)

var serviceSet = wire.NewSet(
	serviceimplement.NewUserService,
)

var repositorySet = wire.NewSet(
	repositoryimplement.NewUserRepository,
)

func InitializeContainer(db database.Db) *controller.ApiContainer {
	wire.Build(serverSet, handlerSet, serviceSet, repositorySet, container)
	return &controller.ApiContainer{}
}
