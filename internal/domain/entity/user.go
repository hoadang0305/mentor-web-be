package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

// mapping for database

type User struct {
	Id    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Role  uint8              `json:"role"`
	Name  string             `json:"name"`
	Email string             `json:"email"`
}
