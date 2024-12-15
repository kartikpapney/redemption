package services

import (
	"context"
	"errors"
	db "redemption/models/db"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateExercise(user primitive.ObjectID, name string, primary []primitive.ObjectID, secondary []primitive.ObjectID, other []primitive.ObjectID) (*db.Exercise, error) {

	exercise := db.NewExercise(user, name, primary, secondary, other)
	err := mgm.Coll(exercise).Create(exercise)
	if err != nil {
		return nil, errors.New("cannot create new user")
	}

	return exercise, nil
}

func UpdateExerciseById(user primitive.ObjectID, exerciseId primitive.ObjectID, name string, primary []primitive.ObjectID, secondary []primitive.ObjectID, other []primitive.ObjectID) (*db.Exercise, error) {
	exercise := &db.Exercise{}
	update := bson.M{
		"$set": bson.M{
			"name":      name,
			"primary":   primary,
			"secondary": secondary,
			"other":     other,
		},
	}
	err := mgm.Coll(exercise).FindOneAndUpdate(context.Background(), exercise, update)
	if err != nil {
		return nil, errors.New("cannot find user")
	}

	return exercise, nil
}

func GetAllExercise(user primitive.ObjectID, limit int64, offset int64) ([]*db.Exercise, int64, error) {
	exercise := []*db.Exercise{}
	opts := &options.FindOptions{
		Limit: &limit,
		Skip:  &offset,
	}

	total, err := mgm.Coll(&db.Exercise{}).CountDocuments(context.Background(), bson.M{"user": user})
	if err != nil {
		return nil, -1, errors.New("failed to count products: " + err.Error())
	}

	err = mgm.Coll(&db.Exercise{}).SimpleFind(&exercise, bson.M{"user": user}, opts)
	if err != nil {
		return nil, -1, errors.New("cannot find exercise")
	}

	return exercise, total, nil
}

func GetExercise(user primitive.ObjectID, exerciseId primitive.ObjectID) (*db.Exercise, error) {
	exercise := db.Exercise{}

	err := mgm.Coll(&db.Exercise{}).FindOne(context.Background(), bson.M{"user": user, "_id": exerciseId}).Decode(&exercise)
	if err != nil {
		return nil, errors.New("cannot find exercise")
	}

	return &exercise, nil
}

func DeleteExerciseById(user primitive.ObjectID, exercise primitive.ObjectID) error {
	delete := bson.M{
		"_id":  exercise,
		"user": user,
	}
	_, err := mgm.Coll(&db.Exercise{}).DeleteOne(context.Background(), delete)
	if err != nil {
		return errors.New("cannot delete exercise")
	}

	return nil
}
