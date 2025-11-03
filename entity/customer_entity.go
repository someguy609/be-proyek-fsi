package entity

// import (
// 	"time"

// 	"github.com/google/uuid"
// )

// type Gender string

// const (
// 	Male   Gender = "M"
// 	Female Gender = "F"
// )

// type Customer struct {
// 	ID        uuid.UUID   `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
// 	Gender    Gender      `gorm:"type:varchar(1);not null" json:"gender" validate:"required"`
// 	EntryTime time.Time   `gorm:"not null" json:"entry_time" validate:"required"`
// 	ExitTime  time.Time   `json:"exit_time"`
// 	Location  *[]Location `gorm:"many2many:customer_locations"`
// }
