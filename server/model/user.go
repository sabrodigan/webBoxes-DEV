package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name     string             `bson:"name, omitempty" json:"name"`
	Email    string             `bson:"email,omitempty" json:"email"`
	Password primitive.ObjectID `bson:"password" json:"password"`
}

type UserCreateDto struct {
	Id       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name     string             `bson:"name, omitempty" json:"name"`
	Email    string             `bson:"email,omitempty" json:"email"`
	Password primitive.ObjectID `bson:"password" json:"password"`
}
