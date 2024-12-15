package services

import (
	"context"
	"errors"
	"fmt"
	db "redemption/models/db"
	requestModel "redemption/models/request"
	"time"

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

	user := db.NewUser(name, email, db.RoleUser, string(password))
	err = mgm.Coll(user).Create(user)

	if err != nil {
		return nil, errors.New("cannot create new user")
	}

	return user, nil
}

func FindUserById(userId primitive.ObjectID) (*db.User, error) {
	user := &db.User{
		Id: userId,
	}

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

func UpdateUserById(userId primitive.ObjectID, update requestModel.UserUpdateRequest) error {
	user := &db.User{}
	updateQuery := bson.M{
		"$set": bson.M{
			"dob":                update.DOB,
			"preferred_unit":     update.PreferredUnit,
			"preferred_exercise": update.PreferredExercise,
			"location":           update.Location,
		},
	}
	_, err := mgm.Coll(user).UpdateByID(context.Background(), userId, updateQuery)
	if err != nil {
		return errors.New("cannot update user")
	}

	return nil
}

func UpdateMeasurementById(userId primitive.ObjectID, preferredUnit db.PreferredUnit, update requestModel.UserMeasurementTrackRequest) error {

	updateQuery := bson.M{
		"$push": bson.M{
			"height": db.HeightEntry{
				Date:   time.Now(),
				Height: update.Height,
				Unit:   preferredUnit.Height,
			},
			"weight": db.WeightEntry{
				Date:   time.Now(),
				Weight: update.Weight,
				Unit:   preferredUnit.Weight,
			},
		},
	}

	_, err := mgm.Coll(&db.User{}).UpdateByID(context.Background(), userId, updateQuery)
	if err != nil {
		return errors.New("cannot update user")
	}

	return nil
}
