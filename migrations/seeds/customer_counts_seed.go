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

func ListCustomerCountSeeder(db *mongo.Database) error {
	jsonFile, err := os.Open("./migrations/json/customer_count.json")
	if err != nil {
		return err
	}

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	var listCustomerCount []entity.CustomerCount
	if err := bson.UnmarshalExtJSON(jsonData, false, &listCustomerCount); err != nil {
		return err
	}

	collection := db.Collection(entity.CustomerCountsCollection)

	for _, data := range listCustomerCount {
		count, err := collection.CountDocuments(context.TODO(), bson.M{"timestamp": data.Timestamp, "location_id": data.LocationID, "gender": data.Gender})
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
