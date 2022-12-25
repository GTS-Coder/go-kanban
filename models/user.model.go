package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 👈 SignUpInput struct
type SignUpInput struct {
	Name         string    `json:"name" bson:"name" binding:"required"`
	Email        string    `json:"email" bson:"email" binding:"required"`
	Password     string    `json:"password" bson:"password" binding:"required,min=6"`
	Role         string    `json:"role" bson:"role"`
	IdGetKanban  string    `json:"id_getkanban" bson:"id_getkanban"`
	AccessToken  string    `json:"accessToken" bson:"accessToken"`
	RefreshToken string    `json:"refreshToken" bson:"refreshToken"`
	Online       time.Time `json:"online" bson:"online"`
}

// 👈 SignInInput struct
type SignInInput struct {
	Email    string `json:"email" bson:"email" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
}

// 👈 GetUSerInput struct using check user exist on server

type GetUSerInput struct {
	Email string `json:"email" bson:"email" binding:"required"`
}

type UserElement struct {
	Name   string    `json:"name" bson:"name" binding:"required"`
	Email  string    `json:"email" bson:"email" binding:"required"`
	Online time.Time `json:"online" bson:"online"`
}

// 👈 DBResponse struct
type DBResponse struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	Name         string             `json:"name" bson:"name"`
	Email        string             `json:"email" bson:"email"`
	Password     string             `json:"password" bson:"password"`
	Role         string             `json:"role" bson:"role"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at"`
	IdGetKanban  string             `json:"id_getkanban" bson:"id_getkanban"`
	AccessToken  string             `json:"accessToken" bson:"accessToken"`
	RefreshToken string             `json:"refreshToken" bson:"refreshToken"`
}

// 👈 Email struct

type Email struct {
	Email string `json:"email" bson:"email" binding:"required"`
}
