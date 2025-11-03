package entity

import (
	"time"
)

type Timestamp struct {
	CreatedAt time.Time `gorm:"type:timestamp with time zone" bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp with time zone" bson:"updated_at" json:"updated_at"`
}

type Authorization struct {
	Token string `json:"token" binding:"required"`
	Role  string `json:"role" binding:"required,oneof=user admin"`
}
