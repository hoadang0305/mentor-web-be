package model

// mapping for request, response

type UserRequest struct {
	Role  int    `json:"role" binding:"required,role"`
	Name  string `json:"name" binding:"required,min=1,max=255"`
	Email string `json:"email" binding:"required,min=1,max=255"`
}
