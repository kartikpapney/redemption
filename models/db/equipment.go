package models

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Equipment struct {
	mgm.DefaultModel `bson:",inline"`
	Id               primitive.ObjectID `form:"id" json:"id" bson:"_id"`
	User             primitive.ObjectID `json:"-" bson:"user"`
	Status           string             `json:"-" bson:"status"`
	Name             string             `json:"name" bson:"name"`
	Resistance       float64            `json:"resistance" bson:"resistance"`
	Unit             string             `json:"unit" bson:"unit"`
}

func NewEquipment(user primitive.ObjectID, name string, resistance float64, unit string) *Equipment {
	return &Equipment{
		Id:         primitive.NewObjectID(),
		User:       user,
		Status:     "Active",
		Name:       name,
		Resistance: resistance,
		Unit:       unit,
	}
}

func (model *Equipment) CollectionName() string {
	return "equipment"
}
