package repository

import (
	"context"
	"time"

	"github.com/someguy609/be-proyek-fsi/entity"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	// "gorm.io/gorm"
)

type (
	CustomerCountRepository interface {
		Create(ctx context.Context, customerCount entity.CustomerCount) (entity.CustomerCount, error)
		// GetAllCustomerCountWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.GetAllCustomerCountRepositoryResponse, error)
		GetCustomerCountByLocation(ctx context.Context, locationId string, start *time.Time, end *time.Time, interval string) ([]entity.CustomerCount, error)
		Update(ctx context.Context, customerCount entity.CustomerCount) (entity.CustomerCount, error)
		Delete(ctx context.Context, customerCountId string) error
	}

	customerCountRepository struct {
		db *mongo.Database
	}
)

func NewCustomerCountRepository(db *mongo.Database) CustomerCountRepository {
	return &customerCountRepository{
		db: db,
	}
}

func (r *customerCountRepository) Create(ctx context.Context, customerCount entity.CustomerCount) (entity.CustomerCount, error) {
	collection := r.db.Collection(entity.CustomerCountsCollection)

	if _, err := collection.InsertOne(ctx, customerCount); err != nil {
		return entity.CustomerCount{}, err
	}

	return customerCount, nil
}

// func (r *customerCountRepository) GetCustomerCountById(ctx context.Context, customerCountId string) (entity.CustomerCount, error) {
// 	collection := r.db.Collection(entity.CustomerCountsCollection)

// 	filter := bson.M{"_id": customerCountId}
// 	collection.FindOne(ctx, filter).Decode(&location)

// 	// if err := tx.WithContext(ctx).Model(&entity.CustomerCount{}).Where("id = ?", customerCountId); err != nil {
// 	// 	return entity.CustomerCount{}, nil
// 	// }

// 	return entity.CustomerCount{}, nil
// }

func (r *customerCountRepository) GetCustomerCountByLocation(
	ctx context.Context,
	locationId string,
	start *time.Time,
	end *time.Time,
	interval string,
) ([]entity.CustomerCount, error) {
	collection := r.db.Collection(entity.CustomerCountsCollection)

	locationObjectId, err := bson.ObjectIDFromHex(locationId)
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"location_id": locationObjectId,
	}

	timeFilter := bson.M{}

	if start != nil {
		timeFilter["$gte"] = bson.NewDateTimeFromTime(*start)
	}

	if end != nil {
		timeFilter["$lte"] = bson.NewDateTimeFromTime(*end)
	}

	if len(timeFilter) > 0 {
		filter["timestamp"] = timeFilter
	}

	group := bson.M{
		"_id": bson.M{
			"timestamp": bson.M{
				"$dateTrunc": bson.M{
					"date": "$timestamp",
					"unit": interval,
				},
			},
			"location_id": "$location_id",
			"gender":      "$gender",
		},
		"count": bson.M{
			"$sum": "$count",
		},
	}

	flatten := bson.M{
		"_id":         0,
		"timestamp":   "$_id.timestamp",
		"location_id": "$_id.location_id",
		"gender":      "$_id.gender",
		"count":       1,
	}

	sort := bson.M{"timestamp": 1}

	pipeline := []bson.M{
		{"$match": filter},
		{"$group": group},
		{"$project": flatten},
		{"$sort": sort},
	}

	opts := options.Aggregate().SetAllowDiskUse(true)

	cursor, err := collection.Aggregate(ctx, pipeline, opts)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var customerCounts []entity.CustomerCount

	if err := cursor.All(ctx, &customerCounts); err != nil {
		return nil, err
	}

	return customerCounts, nil
}

func (r *customerCountRepository) Update(ctx context.Context, customerCount entity.CustomerCount) (entity.CustomerCount, error) {
	collection := r.db.Collection(entity.CustomerCountsCollection)

	filter := bson.M{
		"timestamp":   customerCount.Timestamp,
		"location_id": customerCount.LocationID,
		"gender":      customerCount.Gender,
	}
	update := bson.M{"$set": customerCount}

	if _, err := collection.UpdateOne(ctx, filter, update); err != nil {
		return entity.CustomerCount{}, err
	}

	return customerCount, nil
}

func (r *customerCountRepository) Delete(ctx context.Context, customerCountId string) error {
	collection := r.db.Collection(entity.CustomerCountsCollection)

	filter := bson.M{"_id": customerCountId}

	if _, err := collection.DeleteOne(ctx, filter); err != nil {
		return err
	}

	return nil
}
