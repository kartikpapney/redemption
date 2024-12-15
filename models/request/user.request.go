package requestModel

import (
	db "redemption/models/db"

	validation "github.com/go-ozzo/ozzo-validation"
)

type UserUpdateRequest struct {
	DOB               string           `json:"dob" bson:"dob"`
	PreferredUnit     db.PreferredUnit `json:"preferred_unit" bson:"preferred_unit"`
	PreferredExercise string           `json:"preferred_exercise" bson:"preferred_exercise"`
	Location          string           `json:"location" bson:"location"`
}

func (a UserUpdateRequest) Validate() error {
	err := validation.ValidateStruct(&a,
		validation.Field(&a.DOB, validation.Required),
		validation.Field(&a.Location, validation.Required),
		validation.Field(&a.PreferredExercise, validation.Required, validation.In("calisthenics", "power-lifting", "body-building", "cross-fit", "yoga")),
	)
	if err != nil {
		return err
	}
	return validation.ValidateStruct(&a.PreferredUnit,
		validation.Field(&a.PreferredUnit.Height, validation.Required, validation.In("cm")),
		validation.Field(&a.PreferredUnit.Weight, validation.Required, validation.In("kg", "lbs")),
	)
}

type UserMeasurementTrackRequest struct {
	Height float64 `json:"height" bson:"height"`
	Weight float64 `json:"weight" bson:"weight"`
}

func (a UserMeasurementTrackRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Height, validation.Required),
		validation.Field(&a.Weight, validation.Required),
	)
}
