package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Exercise struct {
	id   primitive.ObjectID `bson:"_id,omitempty"`
	name string             `bson:"name,omitempty"`
}
