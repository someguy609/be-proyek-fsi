package entity

import (
	"time"

	"github.com/google/uuid"
)

type Gender string

const (
	Male   Gender = "M"
	Female Gender = "F"
)

type CustomerCount struct {
	Timestamp  time.Time `gorm:"primaryKey;not null" json:"timestamp" validate:"required"`
	LocationID uuid.UUID `gorm:"type:uuid;primaryKey;not null" json:"location_id" validate:"required,uuid"`
	Location   Location  `gorm:"foreignKey:LocationID;references:ID" json:"-"`
	Gender     Gender    `gorm:"type:varchar(1);primaryKey;not null" json:"gender" validate:"required"`
	Count      uint64    `gorm:"not null" json:"count" validate:"required"`
}
