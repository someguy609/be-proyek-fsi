package entity

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Gender string

const (
	CustomerCountsCollection        = "customer_counts"
	Male                     Gender = "M"
	Female                   Gender = "F"
)

type CustomerCount struct {
	ID         bson.ObjectID `bson:"_id,omitempty" json:"-"`
	LocationID bson.ObjectID `bson:"location_id,omitempty" json:"location_id,omitempty" validate:"required"`
	Timestamp  bson.DateTime `bson:"timestamp" json:"timestamp" validate:"required"`
	Gender     Gender        `bson:"gender" json:"gender" validate:"required"`
	Count      int64         `bson:"count" json:"count" validate:"required"`
}
