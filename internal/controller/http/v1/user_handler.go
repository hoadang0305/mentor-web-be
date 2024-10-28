package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/hoadang0305/mentor-web-be/internal/domain/entity"
	"github.com/hoadang0305/mentor-web-be/internal/domain/http_common"
	"github.com/hoadang0305/mentor-web-be/internal/domain/model"
	"github.com/hoadang0305/mentor-web-be/internal/service"
	"github.com/hoadang0305/mentor-web-be/internal/utils/validation"
	"net/http"
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
	users := handler.userService.GetAllUser(c.Request.Context())
	c.JSON(200, users)
}

// @BasePath /api/v1
// @Summary Add new users
// @Description Add new users
// @Tags Users
// @Produce  json
// @Param params body model.UserRequest true "User payload"
// @Router /api/v1/user [post]
// @Success 200 {object} entity.User
func (handler *UserHandler) Add(c *gin.Context) {
	var userReq model.UserRequest

	if err := validation.BindJsonAndValidate(c, &userReq); err != nil {
		return
	}

	user, err := handler.userService.CreateUser(c.Request.Context(), userReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, http_common.NewErrorResponse(http_common.Error{
			Message: err.Error(), Field: "", Code: http_common.ErrorResponseCode.InternalServerError,
		}))
		return
	}
	c.JSON(http.StatusCreated, http_common.NewSuccessResponse[entity.User](user))
}
