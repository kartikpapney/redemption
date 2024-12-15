package requestModel

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateExerciseRequest struct {
	Name      string               `json:"name" bson:"name"`
	Primary   []primitive.ObjectID `json:"primary" bson:"primary"`
	Secondary []primitive.ObjectID `json:"secondary" bson:"secondary"`
	Other     []primitive.ObjectID `json:"other" bson:"other"`
}

func (a CreateExerciseRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Name, validation.Required),
	)
}

type UpdateExerciseRequest struct {
	Name      string               `json:"name" bson:"name"`
	Primary   []primitive.ObjectID `json:"primary" bson:"primary"`
	Secondary []primitive.ObjectID `json:"secondary" bson:"secondary"`
	Other     []primitive.ObjectID `json:"other" bson:"other"`
}

func (a UpdateExerciseRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Name, validation.Required),
	)
}
