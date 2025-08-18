package migrations

import (
	"github.com/someguy609/be-proyek-fsi/migrations/seeds"
	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) error {
	if err := seeds.ListUserSeeder(db); err != nil {
		return err
	}

	if err := seeds.ListLocationSeeder(db); err != nil {
		return err
	}

	if err := seeds.ListCustomerCountSeeder(db); err != nil {
		return err
	}

	return nil
}
