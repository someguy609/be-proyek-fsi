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
		Create(ctx context.Context, req dto.CustomerCountCreateRequest) (dto.CustomerCountResponse, error)
		GetCustomerCountById(ctx context.Context, customerCountId uint64, start *time.Time, end *time.Time, interval *time.Duration) (dto.CustomerCountGetResponse, error)
		// Update(ctx context.Context, req dto.CustomerCountUpdateRequest, customerCountId uint64) (dto.CustomerCountUpdateResponse, error)
		// Delete(ctx context.Context, customerCountId uint64) error
	}

	customerCountService struct {
		customerCountRepo repository.CustomerCountRepository
		refreshTokenRepo  repository.RefreshTokenRepository
		jwtService        JWTService
		db                *gorm.DB
	}
)

func NewCustomerCountService(
	customerCountRepo repository.CustomerCountRepository,
	db *gorm.DB,
) CustomerCountService {
	return &customerCountService{
		customerCountRepo: customerCountRepo,
		db:                db,
	}
}

func (s *customerCountService) Create(ctx context.Context, req dto.CustomerCountCreateRequest) (dto.CustomerCountResponse, error) {
	customerCount := entity.CustomerCount{}

	customerCountReg, err := s.customerCountRepo.Create(ctx, nil, customerCount)
	if err != nil {
		return dto.CustomerCountResponse{}, dto.ErrCreateCustomerCount
	}

	return dto.CustomerCountResponse{
		ID:         customerCountReg.ID,
		LocationID: customerCountReg.LocationID,
		Count:      customerCountReg.Count,
		Timestamp:  customerCountReg.Timestamp,
	}, nil
}

func (s *customerCountService) GetCustomerCountById(
	ctx context.Context,
	customerCountId uint64,
	start *time.Time,
	end *time.Time,
	interval *time.Duration,
) (dto.CustomerCountGetResponse, error) {
	customerCounts, err := s.customerCountRepo.GetCustomerCountById(ctx, nil, customerCountId, start, end, interval)
	if err != nil {
		return dto.CustomerCountGetResponse{}, dto.ErrGetCustomerCountById
	}

	return dto.CustomerCountGetResponse{
		Data: customerCounts,
	}, nil
}

// func (s *customerCountService) Update(ctx context.Context, req dto.CustomerCountUpdateRequest, customerCountId uint64) (
// 	dto.CustomerCountUpdateResponse,
// 	error,
// ) {
// 	customerCount, err := s.customerCountRepo.GetCustomerCountById(ctx, nil, customerCountId)
// 	if err != nil {
// 		return dto.CustomerCountUpdateResponse{}, dto.ErrCustomerCountNotFound
// 	}

// 	data := entity.CustomerCount{
// 		ID:         customerCount.ID,
// 		LocationID: req.LocationID,
// 		Count:      req.Count,
// 		Timestamp:  req.Timestamp,
// 	}

// 	customerCountUpdate, err := s.customerCountRepo.Update(ctx, nil, data)
// 	if err != nil {
// 		return dto.CustomerCountUpdateResponse{}, dto.ErrUpdateCustomerCount
// 	}

// 	return dto.CustomerCountUpdateResponse{
// 		ID:        customerCountUpdate.ID,
// 		ZoneID:    customerCountUpdate.ZoneID,
// 		Count:     customerCountUpdate.Count,
// 		Timestamp: customerCountUpdate.Timestamp,
// 	}, nil
// }

// func (s *customerCountService) Delete(ctx context.Context, customerCountId uint64) error {
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
