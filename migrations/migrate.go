package migrations

import (
	"github.com/someguy609/be-proyek-fsi/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&entity.User{},
		&entity.RefreshToken{},
		&entity.Location{},
		&entity.CustomerCount{},
	); err != nil {
		return err
	}

	return nil
}
