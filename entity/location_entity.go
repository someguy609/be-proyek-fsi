package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Location struct {
	ID       uuid.UUID       `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	CameraID uint            `gorm:"autoIncrement" json:"camera_id" validate:"required"`
	Name     string          `gorm:"unique;not null" json:"name" validate:"required"`
	X1       float32         `gorm:"not null" json:"x1" validate:"required"`
	Y1       float32         `gorm:"not null" json:"y1" validate:"required"`
	X2       float32         `gorm:"not null" json:"x2" validate:"required"`
	Y2       float32         `gorm:"not null" json:"y2" validate:"required"`
	Counts   []CustomerCount `gorm:"foreignKey:LocationID" json:"-"`

	Timestamp
}

func (l *Location) BeforeCreate(_ *gorm.DB) (err error) {
	if l.ID == uuid.Nil {
		l.ID = uuid.New()
	}

	return nil
}
