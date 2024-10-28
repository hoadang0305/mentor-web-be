package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

// mapping for database

type User struct {
	Id    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Role  int                `json:"role"`
	Name  string             `json:"name"`
	Email string             `json:"email"`
}
