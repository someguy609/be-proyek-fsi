package entity

import "time"

type CustomerCount struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Count      uint64    `gorm:"type:numeric;not null" json:"count" validate:"required"`
	LocationID uint64    `gorm:"index;not null" validate:"required"`
	Timestamp  time.Time `gorm:"index;not null" validate:"required"`
}
