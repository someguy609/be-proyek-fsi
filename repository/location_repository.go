package repository

import (
	"context"

	"github.com/someguy609/be-proyek-fsi/dto"
	"github.com/someguy609/be-proyek-fsi/entity"
	"gorm.io/gorm"
)

type (
	LocationRepository interface {
		Create(ctx context.Context, tx *gorm.DB, location entity.Location) (entity.Location, error)
		GetAllLocationWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.GetAllLocationRepositoryResponse, error)
		GetLocationById(ctx context.Context, tx *gorm.DB, locationId uint64) (entity.Location, error)
		Update(ctx context.Context, tx *gorm.DB, location entity.Location) (entity.Location, error)
		Delete(ctx context.Context, tx *gorm.DB, locationId uint64) error
	}

	locationRepository struct {
		db *gorm.DB
	}
)

func NewLocationRepository(db *gorm.DB) LocationRepository {
	return &locationRepository{
		db: db,
	}
}

func (r *locationRepository) Create(ctx context.Context, tx *gorm.DB, location entity.Location) (entity.Location, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&location).Error; err != nil {
		return entity.Location{}, err
	}

	return location, nil
}

func (r *locationRepository) GetAllLocationWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.GetAllLocationRepositoryResponse, error) {
	if tx == nil {
		tx = r.db
	}

	var locations []entity.Location
	var count int64

	req.Default()

	query := tx.WithContext(ctx).Model(&entity.Location{})

	if req.Search != "" {
		query = query.Where("name LIKE ?", "%"+req.Search+"%")
	}

	if err := query.Count(&count).Error; err != nil {
		return dto.GetAllLocationRepositoryResponse{}, err
	}

	if err := query.Scopes(Paginate(req)).Find(&locations).Error; err != nil {
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

func (r *locationRepository) GetLocationById(ctx context.Context, tx *gorm.DB, locationId uint64) (entity.Location, error) {
	if tx == nil {
		tx = r.db
	}

	var location entity.Location

	if err := tx.WithContext(ctx).Where("id = ?", locationId).Take(&location).Error; err != nil {
		return entity.Location{}, err
	}

	return location, nil
}

func (r *locationRepository) Update(ctx context.Context, tx *gorm.DB, location entity.Location) (entity.Location, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Updates(&location).Error; err != nil {
		return entity.Location{}, err
	}

	return location, nil
}

func (r *locationRepository) Delete(ctx context.Context, tx *gorm.DB, locationId uint64) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Delete(&entity.Location{}, "id = ?", locationId).Error; err != nil {
		return err
	}

	return nil
}
