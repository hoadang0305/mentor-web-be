package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/hoadang0305/mentor-web-be/internal/service"
)

type StudentHandler struct {
	studentService service.StudentService
}

func NewStudentHandler(studentService service.StudentService) *StudentHandler {
	return &StudentHandler{studentService: studentService}
}

// @BasePath /api/v1
// @Summary Get all students
// @Description Get all students
// @Tags Student
// @Produce  json
// @Router /api/v1/students [get]
// @Success 200 {object} []model.Student
func (handler *StudentHandler) GetAll(c *gin.Context) {
	print(123)
	students := handler.studentService.GetAllStudent(c.Request.Context())
	c.JSON(200, students)
}
