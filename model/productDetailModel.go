package model

import (
	"time"
	//"github.com/google/uuid"
	"github.com/pharma999/vender/enum"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type ProductDetail struct {
	ID       		bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	VenderID 		bson.ObjectID `json:"vender_id" bson:"vender_id"`
	//ProductID       *uuid.UUID     `json:"product_id" bson:"product_id" validate:"required"`
	FirstName       string        `json:"first_name" bson:"first_name" validate:"required,min=2,max=100"`
	LastName        string        `json:"last_name" bson:"last_name" validate:"required,min=2,max=100"`
	Email           string        `json:"email" bson:"email" validate:"required,email"`
	PhoneNumber     int64         `json:"phone_number" bson:"phone_number" validate:"required"`
	CreatedAt       time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at" bson:"updated_at"`
	Token           string        `json:"token" bson:"token"`
	RefreshToken    string        `json:"refresh_token" bson:"refresh_token"`
	ProductType     enum.ProductType `json:"product_type" bson:"product_type" validate:"required,oneof=QUICK SCHEDULE EMERGENCY"`
}