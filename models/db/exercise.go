package models

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Exercise struct {
	mgm.DefaultModel `bson:",inline"`
	Id               primitive.ObjectID   `form:"id" json:"id" bson:"_id"`
	User             primitive.ObjectID   `json:"-" bson:"user"`
	Status           string               `json:"-" bson:"status"`
	Name             string               `json:"name" bson:"name"`
	Primary          []primitive.ObjectID `json:"primary" bson:"primary"`
	Secondary        []primitive.ObjectID `json:"secondary" bson:"secondary"`
	Other            []primitive.ObjectID `json:"other" bson:"other"`
}

func NewExercise(user primitive.ObjectID, name string, primary []primitive.ObjectID, secondary []primitive.ObjectID, other []primitive.ObjectID) *Exercise {
	return &Exercise{
		Id:        primitive.NewObjectID(),
		User:      user,
		Status:    "Active",
		Name:      name,
		Primary:   primary,
		Secondary: secondary,
		Other:     other,
	}
}

func (model *Exercise) CollectionName() string {
	return "exercise"
}
