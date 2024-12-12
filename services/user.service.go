package services

import (
	"context"
	"errors"
	"fmt"
	db "redemption/models/db"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(name string, email string, plainPassword string, location string) (*db.User, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("cannot generate hashed password")
	}

	user := db.NewUser(email, string(password), name, db.RoleUser, location)
	err = mgm.Coll(user).Create(user)
	if err != nil {
		return nil, errors.New("cannot create new user")
	}

	return user, nil
}

func FindUserById(userId primitive.ObjectID) (*db.User, error) {
	user := &db.User{}
	err := mgm.Coll(user).FindByID(userId, user)
	if err != nil {
		return nil, errors.New("cannot find user")
	}

	return user, nil
}

func FindUserByEmail(email string) (*db.User, error) {
	user := &db.User{}
	err := mgm.Coll(user).First(bson.M{"email": email}, user)
	if err != nil {
		return nil, errors.New("cannot find user")
	} else if !user.MailVerified {
		return nil, errors.New("email is not verified")
	}

	return user, nil
}

func CheckUserMail(email string) error {
	user := &db.User{}
	userCollection := mgm.Coll(user)
	err := userCollection.First(bson.M{"email": email}, user)
	if err == nil {
		return errors.New("email is already in use")
	}

	return nil
}

func VerifyUserEmail(userId primitive.ObjectID) (*db.User, error) {
	
	user := &db.User{}
	userCollection := mgm.Coll(user)
	context := context.Background()
	err := userCollection.FindByID(userId, user)

	if err != nil {
		return nil, errors.New("user not found")
	} else if user.MailVerified {
		return nil, errors.New("email is already verified")
	}

	update := bson.M{
		"$set": bson.M{"mail_verified": true},
	}

	err = userCollection.FindOneAndUpdate(context, bson.M{"_id": userId}, update).Decode(&user)
	if err != nil {
		return user, errors.New(fmt.Sprintf("failed to update user: %v", err))
	}

	return user, nil
}

