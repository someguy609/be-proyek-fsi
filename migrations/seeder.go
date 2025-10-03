package migrations

import (
	"github.com/someguy609/be-proyek-fsi/migrations/seeds"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func Seeder(db *mongo.Database) error {
	// if err := seeds.ListUserSeeder(db); err != nil {
	// 	return err
	// }

	if err := seeds.ListLocationSeeder(db); err != nil {
		return err
	}

	if err := seeds.ListCustomerCountSeeder(db); err != nil {
		return err
	}

	return nil
}
