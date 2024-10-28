package startup

import (
	"github.com/hoadang0305/mentor-web-be/internal"
	"github.com/hoadang0305/mentor-web-be/internal/controller"
	"github.com/hoadang0305/mentor-web-be/internal/database"
	"github.com/hoadang0305/mentor-web-be/internal/utils/validation"
)

func registerDependencies() *controller.ApiContainer {
	// Open database connection
	db := database.ConnectDB()

	return internal.InitializeContainer(db)
}

func Execute() {
	validation.GetValidations()
	container := registerDependencies()
	container.HttpServer.Run()
}
