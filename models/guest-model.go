package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Guest struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Name     string             `json:"name,omitempty" validate:"required"`
	Surname  string             `json:"surname,omitempty" validate:"required"`
	IsActive *bool              `json:"isActive" binding:"exists"`
	Message  string             `json:"message,omitempty"`
}
