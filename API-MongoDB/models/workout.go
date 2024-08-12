package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Workout struct {
	id        primitive.ObjectID   `bson:"_id,omitempty"`
	exercises []primitive.ObjectID `bson:"exercises"`
	date      primitive.DateTime   `bson:"date,omitempty"`
}
