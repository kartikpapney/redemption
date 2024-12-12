package models

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/kamva/mgm/v3"
)

const (
	RoleUser = "user"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Location		 string `json:"location" bson:"location"`
	Email            string `json:"email" bson:"email"`
	Password         string `json:"-" bson:"password"`
	Role             string `json:"role" bson:"role"`
	MailVerified     bool   `json:"-" bson:"mail_verified"`
}

type UserClaims struct {
	jwt.RegisteredClaims
	Email string `json:"email"`
	Type  string `json:"type"`
}

func NewUser(email string, password string, name string, role string, location string) *User {
	return &User{
		Email:        email,
		Password:     password,
		Name:         name,
		Role:         role,
		Location:     location,
		MailVerified: false,
	}
}

func (model *User) CollectionName() string {
	return "users"
}