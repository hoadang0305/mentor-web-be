package model

// mapping for request, response

type UserRequest struct {
	Role  uint8  `json:"role"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
