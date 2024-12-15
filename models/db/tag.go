package models

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tag struct {
	mgm.DefaultModel `bson:",inline"`
	Id               primitive.ObjectID `form:"id" json:"id" bson:"_id"`
	User             primitive.ObjectID `json:"-" bson:"user"`
	Status           string             `json:"-" bson:"status"`
	Name             string             `json:"name" bson:"name"`
}

func NewTag(user primitive.ObjectID, name string) *Tag {
	return &Tag{
		Id:     primitive.NewObjectID(),
		User:   user,
		Status: "Active",
		Name:   name,
	}
}

func (model *Tag) CollectionName() string {
	return "tag"
}
