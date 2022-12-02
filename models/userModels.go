package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `json:"id"`
	Name          *string            `json:"displayName" validate:"required,min=2,max=10"`
	Email         *string            `json:"email" validate:"required,email"`
	Password      *string            `json:"password"`
	Token         *string            `json:"accessToken"`
	Refresh_token *string            `json:"refresh_token"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
	User_id       string             `json:"user_id"` //related SQL
	Role          string             `json:"role"`
	PhotoUrl      string             `json:"photoUrl"`
}
