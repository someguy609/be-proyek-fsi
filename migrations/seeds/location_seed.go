package seeds

import (
	"context"
	// "encoding/json"
	"io"
	"os"

	"github.com/someguy609/be-proyek-fsi/entity"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func ListLocationSeeder(db *mongo.Database) error {
	jsonFile, err := os.Open("./migrations/json/locations.json")
	if err != nil {
		return err
	}

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	var listLocation []entity.Location
	if err := bson.UnmarshalExtJSON(jsonData, false, &listLocation); err != nil {
		return err
	}

	collection := db.Collection(entity.LocationCollection)

	for _, data := range listLocation {
		count, err := collection.CountDocuments(context.TODO(), bson.M{"name": data.Name})
		if err != nil {
			return err
		}

		if count > 0 {
			continue
		}

		if _, err := collection.InsertOne(context.TODO(), data); err != nil {
			return err
		}
	}

	return nil
}
