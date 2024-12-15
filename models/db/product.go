package models

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	mgm.DefaultModel `bson:",inline"`
	Name         string `json:"name"`
	Manufacturer string `json:"manufacturer"`
	CreatedBy    primitive.ObjectID `json:"created_by" bson:"created_by"`
}

func NewProduct(name string, manufacturer string, createdBy primitive.ObjectID) *Product {
	return &Product{
		Name:         name,
		Manufacturer: manufacturer,
		CreatedBy: createdBy,
	}
}

func (model *Product) CollectionName() string {
	return "product"
}