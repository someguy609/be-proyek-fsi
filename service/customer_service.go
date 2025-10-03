package service

// import (
// 	"context"
// 	"time"

// 	"gorm.io/gorm"

// 	"github.com/someguy609/be-proyek-fsi/dto"
// 	"github.com/someguy609/be-proyek-fsi/entity"
// 	"github.com/someguy609/be-proyek-fsi/repository"
// )

// type (
// 	CustomerService interface {
// 		Create(ctx context.Context, req dto.CustomerCreateRequest) (dto.CustomerResponse, error)
// 		GetCustomerById(ctx context.Context, customerId string) (dto.CustomerResponse, error)
// 		GetCustomerCountByLocation(ctx context.Context, locationId string, start *time.Time, end *time.Time, interval string) (dto.CustomerCountResponse, error)
// 		Update(ctx context.Context, req dto.CustomerUpdateRequest, locationId string) (dto.CustomerUpdateResponse, error)
// 		Delete(ctx context.Context, customerId string) error
// 	}

// 	customerService struct {
// 		customerRepo repository.CustomerRepository
// 		locationRepo repository.LocationRepository
// 		db           *gorm.DB
// 	}
// )

// func NewCustomerService(
// 	customerRepo repository.CustomerRepository,
// 	locationRepo repository.LocationRepository,
// 	db *gorm.DB,
// ) CustomerService {
// 	return &customerService{
// 		customerRepo: customerRepo,
// 		locationRepo: locationRepo,
// 		db:           db,
// 	}
// }

// func (s *customerService) Create(ctx context.Context, req dto.CustomerCreateRequest) (dto.CustomerResponse, error) {
// 	location, err := s.locationRepo.GetLocationById(ctx, nil, req.LocationID)

// 	if err != nil {
// 		return dto.CustomerResponse{}, dto.ErrLocationNotFound
// 	}

// 	customer := entity.Customer{
// 		Gender:    req.Gender,
// 		EntryTime: req.EntryTime,
// 		ExitTime:  req.ExitTime,
// 	}

// 	customerReg, err := s.customerRepo.Create(ctx, nil, customer, location)
// 	if err != nil {
// 		return dto.CustomerResponse{}, dto.ErrCreateCustomer
// 	}

// 	return dto.CustomerResponse{
// 		// Timestamp:  customerReg.Timestamp,
// 		// LocationID: customerReg.LocationID.String(),
// 		Gender:    customerReg.Gender,
// 		EntryTime: customerReg.EntryTime,
// 		ExitTime:  customerReg.ExitTime,
// 		// Count:      customerReg.Count,
// 	}, nil
// }

// func (s *customerService) GetCustomerById(
// 	ctx context.Context,
// 	customerId string,
// ) (dto.CustomerResponse, error) {
// 	customer, err := s.customerRepo.GetCustomerById(ctx, nil, customerId)
// 	if err != nil {
// 		return dto.CustomerResponse{}, err
// 	}
// 	return dto.CustomerResponse{
// 		ID:        customer.ID,
// 		Gender:    customer.Gender,
// 		EntryTime: customer.EntryTime,
// 		ExitTime:  customer.ExitTime,
// 	}, nil
// }

// func (s *customerService) GetCustomerCountByLocation(
// 	ctx context.Context,
// 	locationId string,
// 	start *time.Time,
// 	end *time.Time,
// 	interval string,
// ) (dto.CustomerCountResponse, error) {
// 	customers, err := s.customerRepo.GetCustomerCountByLocation(ctx, nil, locationId, start, end, interval)
// 	if err != nil {
// 		return dto.CustomerCountResponse{}, dto.ErrGetCustomerById
// 	}

// 	return dto.CustomerCountResponse{
// 		Data: customers,
// 	}, nil
// }

// func (s *customerService) Update(ctx context.Context, req []dto.CustomerUpdateRequest, locationId string) (
// 	dto.CustomerUpdateResponse,
// 	error,
// ) {
// 	location, err := s.locationRepo.GetLocationById(ctx, nil, locationId)

// 	if err != nil {
// 		return dto.CustomerUpdateResponse{}, dto.ErrLocationNotFound
// 	}

// 	var response dto.CustomerUpdateResponse

// 	for _, point := range req {
// 		data := entity.Customer{
// 			Timestamp:  point.Timestamp,
// 			LocationID: location.ID,
// 			Gender:     point.Gender,
// 			Count:      point.Count,
// 		}

// 		customerUpdate, err := s.customerRepo.Update(ctx, nil, data)

// 		if err != nil {
// 			return dto.CustomerUpdateResponse{}, dto.ErrUpdateCustomer
// 		}

// 		response.Data = append(response.Data, customerUpdate)
// 	}

// 	return response, nil
// }

// // func (s *customerService) Delete(ctx context.Context, locationId string, timestamp *time.Time) error {
// // 	tx := s.db.Begin()
// // 	defer SafeRollback(tx)

// // 	customer, err := s.customerRepo.GetCustomerById(ctx, nil, customerId)
// // 	if err != nil {
// // 		return dto.ErrCustomerNotFound
// // 	}

// // 	err = s.customerRepo.Delete(ctx, nil, customer.ID)
// // 	if err != nil {
// // 		return dto.ErrDeleteCustomer
// // 	}

// // 	return nil
// // }
