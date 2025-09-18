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
		GetLocationById(ctx context.Context, locationId string) (dto.LocationResponse, error)
		Update(ctx context.Context, req dto.LocationUpdateRequest, locationId string) (dto.LocationUpdateResponse, error)
		Delete(ctx context.Context, locationId string) error
	}

	locationService struct {
		locationRepo repository.LocationRepository
		db           *gorm.DB
	}
)

func NewLocationService(
	locationRepo repository.LocationRepository,
	db *gorm.DB,
) LocationService {
	return &locationService{
		locationRepo: locationRepo,
		db:           db,
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
		CameraID: req.CameraID,
		Name: req.Name,
		X1:   req.X1,
		Y1:   req.Y1,
		X2:   req.X2,
		Y2:   req.Y2,
	}

	locationReg, err := s.locationRepo.Create(ctx, nil, location)
	if err != nil {
		return dto.LocationResponse{}, dto.ErrCreateLocation
	}

	return dto.LocationResponse{
		ID:   locationReg.ID.String(),
		CameraID: locationReg.CameraID,
		Name: locationReg.Name,
		X1:   locationReg.X1,
		Y1:   locationReg.Y1,
		X2:   locationReg.X2,
		Y2:   locationReg.Y2,
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
			ID:   location.ID.String(),
			Name: location.Name,
			X1:   location.X1,
			Y1:   location.Y1,
			X2:   location.X2,
			Y2:   location.Y2,
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

func (s *locationService) GetLocationById(ctx context.Context, locationId string) (dto.LocationResponse, error) {
	location, err := s.locationRepo.GetLocationById(ctx, nil, locationId)
	if err != nil {
		return dto.LocationResponse{}, dto.ErrGetLocationById
	}

	return dto.LocationResponse{
		ID:   location.ID.String(),
		Name: location.Name,
		X1:   location.X1,
		Y1:   location.Y1,
		X2:   location.X2,
		Y2:   location.Y2,
	}, nil
}

func (s *locationService) Update(ctx context.Context, req dto.LocationUpdateRequest, locationId string) (
	dto.LocationUpdateResponse,
	error,
) {
	location, err := s.locationRepo.GetLocationById(ctx, nil, locationId)
	if err != nil {
		return dto.LocationUpdateResponse{}, dto.ErrLocationNotFound
	}

	data := entity.Location{
		ID:   location.ID,
		Name: req.Name,
		X1:   req.X1,
		Y1:   req.Y1,
		X2:   req.X2,
		Y2:   req.Y2,
	}

	locationUpdate, err := s.locationRepo.Update(ctx, nil, data)
	if err != nil {
		return dto.LocationUpdateResponse{}, dto.ErrUpdateLocation
	}

	return dto.LocationUpdateResponse{
		ID:   locationUpdate.ID.String(),
		Name: locationUpdate.Name,
		X1:   locationUpdate.X1,
		Y1:   locationUpdate.Y1,
		X2:   locationUpdate.X2,
		Y2:   locationUpdate.Y2,
	}, nil
}

func (s *locationService) Delete(ctx context.Context, locationId string) error {
	tx := s.db.Begin()
	defer SafeRollback(tx)

	location, err := s.locationRepo.GetLocationById(ctx, nil, locationId)
	if err != nil {
		return dto.ErrLocationNotFound
	}

	err = s.locationRepo.Delete(ctx, nil, location.ID.String())
	if err != nil {
		return dto.ErrDeleteLocation
	}

	return nil
}
