package repository

import (
	"context"

	"github.com/someguy609/be-proyek-fsi/dto"
	"github.com/someguy609/be-proyek-fsi/entity"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type (
	LocationRepository interface {
		Create(ctx context.Context, location entity.Location) (entity.Location, error)
		GetAllLocationWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.GetAllLocationRepositoryResponse, error)
		GetLocationById(ctx context.Context, locationId string) (entity.Location, error)
		Update(ctx context.Context, location entity.Location) (entity.Location, error)
		Delete(ctx context.Context, locationId string) error
	}

	locationRepository struct {
		db *mongo.Database
	}
)

func NewLocationRepository(db *mongo.Database) LocationRepository {
	return &locationRepository{
		db: db,
	}
}

func (r *locationRepository) Create(ctx context.Context, location entity.Location) (entity.Location, error) {
	collection := r.db.Collection(entity.LocationCollection)

	res, err := collection.InsertOne(ctx, location)
	if err != nil {
		return entity.Location{}, nil
	}

	location.ID = res.InsertedID.(bson.ObjectID);

	return location, nil
}

func (r *locationRepository) GetAllLocationWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.GetAllLocationRepositoryResponse, error) {
	collection := r.db.Collection(entity.LocationCollection)

	var locations []entity.Location
	var count int64

	req.Default()

	filter := bson.M{}

	if req.Search != "" {
		filter["name"] = bson.M{"$regex": req.Search, "$options": "i"}
	}

	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return dto.GetAllLocationRepositoryResponse{}, err
	}

	findOptions := options.Find().SetSkip(int64((req.Page - 1) * req.PerPage)).SetLimit(int64(req.PerPage))
	cursor, err := collection.Find(ctx, filter, findOptions)

	if err != nil {
		return dto.GetAllLocationRepositoryResponse{}, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &locations); err != nil {
		return dto.GetAllLocationRepositoryResponse{}, err
	}

	totalPage := TotalPage(count, int64(req.PerPage))

	return dto.GetAllLocationRepositoryResponse{
		Locations: locations,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			Count:   count,
			MaxPage: totalPage,
		},
	}, nil
}

func (r *locationRepository) GetLocationById(ctx context.Context, locationId string) (entity.Location, error) {
	collection := r.db.Collection(entity.LocationCollection)

	var location entity.Location

	locationObjectId, err := bson.ObjectIDFromHex(locationId)
	if err != nil {
		return entity.Location{}, err
	}

	filter := bson.M{"_id": locationObjectId}

	if err := collection.FindOne(ctx, filter).Decode(&location); err != nil {
		return entity.Location{}, err
	}

	return location, nil
}

func (r *locationRepository) Update(ctx context.Context, location entity.Location) (entity.Location, error) {
	collection := r.db.Collection(entity.LocationCollection)
	
	filter := bson.M{"_id": location.ID}
	update := bson.M{"$set": location}
	
	if _, err := collection.UpdateOne(ctx, filter, update); err != nil {
		return entity.Location{}, err
	}
	
	return location, nil
}

func (r *locationRepository) Delete(ctx context.Context, locationId string) error {
	collection := r.db.Collection(entity.LocationCollection)
	
	locationObjectId, err := bson.ObjectIDFromHex(locationId)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": locationObjectId}

	if _, err := collection.DeleteOne(ctx, filter); err != nil {
		return err
	}

	return nil
}
