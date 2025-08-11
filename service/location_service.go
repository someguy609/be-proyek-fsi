package service

import (
	"context"

	"gorm.io/gorm"

	"github.com/someguy609/be-proyek-fsi/dto"
	"github.com/someguy609/be-proyek-fsi/entity"
	"github.com/someguy609/be-proyek-fsi/repository"
)

type (
	LocationService interface {
		Create(ctx context.Context, req dto.LocationCreateRequest) (dto.LocationResponse, error)
		GetAllLocationWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.LocationPaginationResponse, error)
		GetLocationById(ctx context.Context, locationId uint64) (dto.LocationResponse, error)
		Update(ctx context.Context, req dto.LocationUpdateRequest, locationId uint64) (dto.LocationUpdateResponse, error)
		Delete(ctx context.Context, locationId uint64) error
	}

	locationService struct {
		locationRepo         repository.LocationRepository
		db               *gorm.DB
	}
)

func NewLocationService(
	locationRepo repository.LocationRepository,
	db *gorm.DB,
) LocationService {
	return &locationService{
		locationRepo: locationRepo,
		db:       db,
	}
}

// func SafeRollback(tx *gorm.DB) {
// 	if r := recover(); r != nil {
// 		tx.Rollback()
// 		// TODO: Do you think that we should panic here?
// 		// panic(r)
// 	}
// }

func (s *locationService) Create(ctx context.Context, req dto.LocationCreateRequest) (dto.LocationResponse, error) {

	location := entity.Location{
		Name:      req.Name,
		Longitude: req.Longitude,
		Latitude:  req.Latitude,
	}

	locationReg, err := s.locationRepo.Create(ctx, nil, location)
	if err != nil {
		return dto.LocationResponse{}, dto.ErrCreateLocation
	}

	return dto.LocationResponse{
		ID:        locationReg.ID,
		Name:      locationReg.Name,
		Longitude: locationReg.Longitude,
		Latitude:  locationReg.Latitude,
	}, nil
}

func (s *locationService) GetAllLocationWithPagination(
	ctx context.Context,
	req dto.PaginationRequest,
) (dto.LocationPaginationResponse, error) {
	dataWithPaginate, err := s.locationRepo.GetAllLocationWithPagination(ctx, nil, req)
	if err != nil {
		return dto.LocationPaginationResponse{}, err
	}

	var datas []dto.LocationResponse
	for _, location := range dataWithPaginate.Locations {
		data := dto.LocationResponse{
			ID:        location.ID,
			Name:      location.Name,
			Longitude: location.Longitude,
			Latitude:  location.Latitude,
		}

		datas = append(datas, data)
	}

	return dto.LocationPaginationResponse{
		Data: datas,
		PaginationResponse: dto.PaginationResponse{
			Page:    dataWithPaginate.Page,
			PerPage: dataWithPaginate.PerPage,
			MaxPage: dataWithPaginate.MaxPage,
			Count:   dataWithPaginate.Count,
		},
	}, nil
}

func (s *locationService) GetLocationById(ctx context.Context, locationId uint64) (dto.LocationResponse, error) {
	location, err := s.locationRepo.GetLocationById(ctx, nil, locationId)
	if err != nil {
		return dto.LocationResponse{}, dto.ErrGetLocationById
	}

	return dto.LocationResponse{
		ID:        location.ID,
		Name:      location.Name,
		Longitude: location.Longitude,
		Latitude:  location.Latitude,
	}, nil
}

func (s *locationService) Update(ctx context.Context, req dto.LocationUpdateRequest, locationId uint64) (
	dto.LocationUpdateResponse,
	error,
) {
	location, err := s.locationRepo.GetLocationById(ctx, nil, locationId)
	if err != nil {
		return dto.LocationUpdateResponse{}, dto.ErrLocationNotFound
	}

	data := entity.Location{
		ID:        location.ID,
		Name:      req.Name,
		Longitude: req.Longitude,
		Latitude:  req.Latitude,
	}

	locationUpdate, err := s.locationRepo.Update(ctx, nil, data)
	if err != nil {
		return dto.LocationUpdateResponse{}, dto.ErrUpdateLocation
	}

	return dto.LocationUpdateResponse{
		ID:        locationUpdate.ID,
		Name:      locationUpdate.Name,
		Longitude: locationUpdate.Longitude,
		Latitude:  locationUpdate.Latitude,
	}, nil
}

func (s *locationService) Delete(ctx context.Context, locationId uint64) error {
	tx := s.db.Begin()
	defer SafeRollback(tx)

	location, err := s.locationRepo.GetLocationById(ctx, nil, locationId)
	if err != nil {
		return dto.ErrLocationNotFound
	}

	err = s.locationRepo.Delete(ctx, nil, location.ID)
	if err != nil {
		return dto.ErrDeleteLocation
	}

	return nil
}
