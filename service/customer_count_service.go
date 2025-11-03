package service

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	// "gorm.io/gorm"

	"github.com/someguy609/be-proyek-fsi/dto"
	"github.com/someguy609/be-proyek-fsi/entity"
	"github.com/someguy609/be-proyek-fsi/repository"
)

type (
	CustomerCountService interface {
		Create(ctx context.Context, req dto.CustomerCountCreateRequest, locationId string) (dto.CustomerCountResponse, error)
		GetCustomerCountByLocation(ctx context.Context, locationId string, start *time.Time, end *time.Time, interval string) (dto.CustomerCountGetResponse, error)
		Update(ctx context.Context, req []dto.CustomerCountUpdateRequest, locationId string) (dto.CustomerCountUpdateResponse, error)
		// Delete(ctx context.Context, locationId string, timestamp *time.Time) error
	}

	customerCountService struct {
		customerCountRepo repository.CustomerCountRepository
		locationRepo      repository.LocationRepository
		db                *mongo.Database
	}
)

func NewCustomerCountService(
	customerCountRepo repository.CustomerCountRepository,
	locationRepo repository.LocationRepository,
	db *mongo.Database,
) CustomerCountService {
	return &customerCountService{
		customerCountRepo: customerCountRepo,
		locationRepo:      locationRepo,
		db:                db,
	}
}

func (s *customerCountService) Create(ctx context.Context, req dto.CustomerCountCreateRequest, locationId string) (dto.CustomerCountResponse, error) {
	location, err := s.locationRepo.GetLocationById(ctx, locationId)

	if err != nil {
		return dto.CustomerCountResponse{}, dto.ErrLocationNotFound
	}

	customerCount := entity.CustomerCount{
		LocationID: location.ID,
		Gender:     req.Gender,
		Count:      req.Count,
		Timestamp:  bson.NewDateTimeFromTime(req.Timestamp),
	}

	customerCountReg, err := s.customerCountRepo.Create(ctx, customerCount)
	if err != nil {
		return dto.CustomerCountResponse{}, dto.ErrCreateCustomerCount
	}

	return dto.CustomerCountResponse{
		Timestamp:  customerCountReg.Timestamp.Time(),
		LocationID: customerCountReg.LocationID.Hex(),
		Gender:     customerCountReg.Gender,
		Count:      customerCountReg.Count,
	}, nil
}

func (s *customerCountService) GetCustomerCountByLocation(
	ctx context.Context,
	locationId string,
	start *time.Time,
	end *time.Time,
	interval string,
) (dto.CustomerCountGetResponse, error) {
	customerCounts, err := s.customerCountRepo.GetCustomerCountByLocation(ctx, locationId, start, end, interval)
	if err != nil {
		return dto.CustomerCountGetResponse{}, dto.ErrGetCustomerCountById
	}

	return dto.CustomerCountGetResponse{
		Data: customerCounts,
	}, nil
}

func (s *customerCountService) Update(ctx context.Context, req []dto.CustomerCountUpdateRequest, locationId string) (
	dto.CustomerCountUpdateResponse,
	error,
) {
	location, err := s.locationRepo.GetLocationById(ctx, locationId)

	if err != nil {
		return dto.CustomerCountUpdateResponse{}, dto.ErrLocationNotFound
	}

	var response dto.CustomerCountUpdateResponse

	for _, point := range req {
		data := entity.CustomerCount{
			Timestamp:  bson.NewDateTimeFromTime(point.Timestamp),
			LocationID: location.ID,
			Gender:     point.Gender,
			Count:      point.Count,
		}

		customerCountUpdate, err := s.customerCountRepo.Update(ctx, data)

		if err != nil {
			return dto.CustomerCountUpdateResponse{}, dto.ErrUpdateCustomerCount
		}

		response.Data = append(response.Data, customerCountUpdate)
	}

	return response, nil
}

// func (s *customerCountService) Delete(ctx context.Context, locationId string, timestamp *time.Time) error {
// 	customerCount, err := s.customerCountRepo.GetCustomerCountById(ctx, nil, customerCountId)
// 	if err != nil {
// 		return dto.ErrCustomerCountNotFound
// 	}

// 	err = s.customerCountRepo.Delete(ctx, nil, customerCount.ID)
// 	if err != nil {
// 		return dto.ErrDeleteCustomerCount
// 	}

// 	return nil
// }
