package startup

import (
	"github.com/hoadang0305/mentor-web-be/internal"
	"github.com/hoadang0305/mentor-web-be/internal/controller"
	"github.com/hoadang0305/mentor-web-be/internal/database"
)

func registerDependencies() *controller.ApiContainer {
	// Open database connection
	db := database.Open()

	return internal.InitializeContainer(db)
}

func Execute() {
	container := registerDependencies()
	container.HttpServer.Run()
}
