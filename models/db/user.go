package models

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	RoleUser = "user"
)

type WeightEntry struct {
	Date   time.Time `json:"date" bson:"date"`
	Weight float64   `json:"weight" bson:"weight"`
	Unit   string    `json:"unit" bson:"unit"`
}

type HeightEntry struct {
	Date   time.Time `json:"date" bson:"date"`
	Height float64   `json:"height" bson:"height"`
	Unit   string    `json:"unit" bson:"unit"`
}

type PreferredUnit struct {
	Height string `json:"height" bson:"height"`
	Weight string `json:"weight" bson:"weight"`
}

type User struct {
	mgm.DefaultModel  `bson:",inline"`
	Id                primitive.ObjectID `form:"id" json:"id" bson:"_id"`
	Status            string             `json:"-" bson:"status"`
	Name              string             `json:"name" bson:"name"`
	Email             string             `json:"email" bson:"email"`
	Password          string             `json:"-" bson:"password"`
	DOB               time.Time          `json:"dob" bson:"dob"`
	PreferredUnit     PreferredUnit      `json:"preferred_unit" bson:"preferred_unit"`
	Height            []HeightEntry      `json:"height" bson:"height"`
	Weight            []WeightEntry      `json:"weight" bson:"weight"`
	PreferredExercise string             `json:"preferred_exercise" bson:"preferred_exercise"`
	Location          string             `json:"location" bson:"location"`
	Role              string             `json:"-" bson:"role"`
	MailVerified      bool               `json:"-" bson:"mail_verified"`
}

type UserClaims struct {
	jwt.RegisteredClaims
	Email string `json:"email"`
	Type  string `json:"type"`
}

func NewUser(name string, email string, role string, password string) *User {
	return &User{
		Id:            primitive.NewObjectID(),
		Status:        "Inactive",
		Name:          name,
		Email:         email,
		Password:      password,
		DOB:           time.Now(),
		Role:          RoleUser,
		MailVerified:  false,
		Height:        []HeightEntry{},
		Weight:        []WeightEntry{},
		PreferredUnit: PreferredUnit{Height: "cm", Weight: "kg"},
	}
}

func (model *User) CollectionName() string {
	return "user"
}
