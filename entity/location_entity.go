package entity

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

const LocationCollection = "locations"

type Location struct {
	ID       bson.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	CameraID int64         `bson:"camera_id" json:"camera_id" validate:"required"`
	Name     string        `bson:"name" json:"name" validate:"required"`
	X1       int64         `bson:"x1" json:"x1" validate:"required"`
	Y1       int64         `bson:"y1" json:"y1" validate:"required"`
	X2       int64         `bson:"x2" json:"x2" validate:"required"`
	Y2       int64         `bson:"y2" json:"y2" validate:"required"`

	Timestamp
}