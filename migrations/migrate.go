package migrations

import (
	"context"

	// "github.com/someguy609/be-proyek-fsi/entity"
	"github.com/someguy609/be-proyek-fsi/entity"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	// "gorm.io/gorm"
)

// func Migrate(db *gorm.DB) error {
// 	if err := db.AutoMigrate(
// 		&entity.User{},
// 		&entity.RefreshToken{},
// 		&entity.Location{},
// 		&entity.CustomerCount{},
// 	); err != nil {
// 		return err
// 	}

// 	return nil
// }

func Migrate(db *mongo.Database) error {
	collections := []struct {
		name    string
		options options.Lister[options.CreateCollectionOptions]
	}{
		{
			name:    entity.LocationCollection,
			options: nil,
		},
		{
			name:    entity.CustomerCountsCollection,
			options: options.CreateCollection().SetTimeSeriesOptions(options.TimeSeries().SetTimeField("timestamp").SetGranularity("minutes")),
		},
	}

	for _, c := range collections {
		if err := db.CreateCollection(context.TODO(), c.name, c.options); err != nil {
			return err
		}
	}

	return nil
}