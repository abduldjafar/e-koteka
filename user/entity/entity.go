package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type CustomerUser struct {
	Name      string `bson:"name" json:"name"`
	Age       int    `bson:"age" json:"age"`
	City      string `bson:"city" json:"city"`
	Email     string `bson:"email" json:"email"`
	Passwords string `bson:"passwords" json:"passwords"`
}

type CustomerUserResponses struct {
	Id        primitive.ObjectID `bson:"_id" json:"_id"`
	Name      string             `bson:"name" json:"name"`
	Age       int                `bson:"age" json:"age"`
	City      string             `bson:"city" json:"city"`
	Email     string             `bson:"email" json:"email"`
	Passwords string             `bson:"passwords" json:"passwords"`
}
