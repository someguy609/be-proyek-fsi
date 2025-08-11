package entity

type Location struct {
	ID        uint64          `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string          `gorm:"type:varchar(255);unique;not null" json:"name" validate:"required"`
	Longitude float32         `gorm:"type:float;not null" json:"longitude" validate:"required"`
	Latitude  float32         `gorm:"type:float;not null" json:"latitude" validate:"required"`

	Timestamp
}
