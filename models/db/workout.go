package models

import (
	"time"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Workout struct {
	mgm.DefaultModel `bson:",inline"`
	Id               primitive.ObjectID `form:"id" json:"id" bson:"_id"`
	User             primitive.ObjectID `json:"-" bson:"user"`
	Exercise         primitive.ObjectID `json:"exercise" bson:"exercise"`
	Equipment        primitive.ObjectID `json:"equipment" bson:"equipment"`
	Timestamp        time.Time          `json:"timestamp" bson:"timestamp"`
	Count            int64              `json:"count" bson:"count"`
	Duration         int64              `json:"duration" bson:"duration"`
}

func NewWorkout(user, exercise, equipment primitive.ObjectID, count, duration int64, timestamp time.Time) *Workout {
	return &Workout{
		Id:        primitive.NewObjectID(),
		User:      user,
		Exercise:  exercise,
		Equipment: equipment,
		Timestamp: timestamp,
		Count:     count,
		Duration:  duration,
	}
}

func (model *Workout) CollectionName() string {
	return "workout"
}
