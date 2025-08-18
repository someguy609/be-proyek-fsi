package service

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/someguy609/be-proyek-fsi/dto"
	"github.com/someguy609/be-proyek-fsi/entity"
	"github.com/someguy609/be-proyek-fsi/repository"
)

type (
	CustomerCountService interface {
		Create(ctx context.Context, req dto.CustomerCountCreateRequest, locationId string) (dto.CustomerCountResponse, error)
		GetCustomerCountByLocation(ctx context.Context, locationId string, start *time.Time, end *time.Time, interval *time.Duration) (dto.CustomerCountGetResponse, error)
		Update(ctx context.Context, req []dto.CustomerCountUpdateRequest, locationId string) (dto.CustomerCountUpdateResponse, error)
		// Delete(ctx context.Context, locationId string, timestamp *time.Time) error
	}

	customerCountService struct {
		customerCountRepo repository.CustomerCountRepository
		locationRepo      repository.LocationRepository
		db                *gorm.DB
	}
)

func NewCustomerCountService(
	customerCountRepo repository.CustomerCountRepository,
	locationRepo repository.LocationRepository,
	db *gorm.DB,
) CustomerCountService {
	return &customerCountService{
		customerCountRepo: customerCountRepo,
		locationRepo:      locationRepo,
		db:                db,
	}
}

func (s *customerCountService) Create(ctx context.Context, req dto.CustomerCountCreateRequest, locationId string) (dto.CustomerCountResponse, error) {
	location, err := s.locationRepo.GetLocationById(ctx, nil, locationId)

	if err != nil {
		return dto.CustomerCountResponse{}, dto.ErrLocationNotFound
	}

	customerCount := entity.CustomerCount{
		LocationID: location.ID,
		Gender:     req.Gender,
		Count:      req.Count,
		Timestamp:  req.Timestamp,
	}

	customerCountReg, err := s.customerCountRepo.Create(ctx, nil, customerCount)
	if err != nil {
		return dto.CustomerCountResponse{}, dto.ErrCreateCustomerCount
	}

	return dto.CustomerCountResponse{
		Timestamp:  customerCountReg.Timestamp,
		LocationID: customerCountReg.LocationID.String(),
		Gender:     customerCountReg.Gender,
		Count:      customerCountReg.Count,
	}, nil
}

func (s *customerCountService) GetCustomerCountByLocation(
	ctx context.Context,
	locationId string,
	start *time.Time,
	end *time.Time,
	interval *time.Duration,
) (dto.CustomerCountGetResponse, error) {
	customerCounts, err := s.customerCountRepo.GetCustomerCountByLocation(ctx, nil, locationId, start, end, interval)
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
	location, err := s.locationRepo.GetLocationById(ctx, nil, locationId)

	if err != nil {
		return dto.CustomerCountUpdateResponse{}, dto.ErrLocationNotFound
	}

	var response dto.CustomerCountUpdateResponse

	for _, point := range req {
		data := entity.CustomerCount{
			Timestamp:  point.Timestamp,
			LocationID: location.ID,
			Gender:     point.Gender,
			Count:      point.Count,
		}

		customerCountUpdate, err := s.customerCountRepo.Update(ctx, nil, data)

		if err != nil {
			return dto.CustomerCountUpdateResponse{}, dto.ErrUpdateCustomerCount
		}

		response.Data = append(response.Data, customerCountUpdate)
	}

	return response, nil
}

// func (s *customerCountService) Delete(ctx context.Context, locationId string, timestamp *time.Time) error {
// 	tx := s.db.Begin()
// 	defer SafeRollback(tx)

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
