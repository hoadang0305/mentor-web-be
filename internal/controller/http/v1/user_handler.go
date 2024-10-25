package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/hoadang0305/mentor-web-be/internal/service"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// @BasePath /api/v1
// @Summary Get all users
// @Description Get all users
// @Tags Users
// @Produce  json
// @Router /api/v1/user [get]
// @Success 200 {object} []entity.User
func (handler *UserHandler) GetAll(c *gin.Context) {
	print(123)
	users := handler.userService.GetAllUser(c.Request.Context())
	c.JSON(200, users)
}
