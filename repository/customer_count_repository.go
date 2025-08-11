package repository

import (
	"context"
	"time"

	"github.com/someguy609/be-proyek-fsi/entity"
	"gorm.io/gorm"
)

type (
	CustomerCountRepository interface {
		Create(ctx context.Context, tx *gorm.DB, customerCount entity.CustomerCount) (entity.CustomerCount, error)
		// GetAllCustomerCountWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.GetAllCustomerCountRepositoryResponse, error)
		GetCustomerCountById(ctx context.Context, tx *gorm.DB, customerCountId uint64, start *time.Time, end *time.Time, interval *time.Duration) ([]entity.CustomerCount, error)
		Update(ctx context.Context, tx *gorm.DB, customerCountId entity.CustomerCount) (entity.CustomerCount, error)
		Delete(ctx context.Context, tx *gorm.DB, customerCountId uint64) error
	}

	customerCountRepository struct {
		db *gorm.DB
	}
)

func NewCustomerCountRepository(db *gorm.DB) CustomerCountRepository {
	return &customerCountRepository{
		db: db,
	}
}

func (r *customerCountRepository) Create(ctx context.Context, tx *gorm.DB, customerCount entity.CustomerCount) (entity.CustomerCount, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&customerCount).Error; err != nil {
		return entity.CustomerCount{}, err
	}

	return customerCount, nil
}

func (r *customerCountRepository) GetCustomerCountById(
	ctx context.Context,
	tx *gorm.DB,
	customerCountId uint64,
	start *time.Time,
	end *time.Time,
	interval *time.Duration,
) ([]entity.CustomerCount, error) {
	if tx == nil {
		tx = r.db
	}

	query := tx.Raw(`
		SELECT
			time_bucket(?, timestamp) AS timestamp
			SUM(count) AS count
		FROM
			customer_count
		WHERE
			location_id = ? AND timestamp >= ? AND timestamp <= ?
		GROUP BY
			time_bucket(?, timestamp)
		ORDER BY
			timestamp

	`, interval, customerCountId, start, end, interval)

	var customerCounts []entity.CustomerCount

	if err := query.Scan(&customerCounts).Error; err != nil {
		return nil, err
	}

	return customerCounts, nil
}

func (r *customerCountRepository) Update(ctx context.Context, tx *gorm.DB, customerCount entity.CustomerCount) (entity.CustomerCount, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Updates(&customerCount).Error; err != nil {
		return entity.CustomerCount{}, err
	}

	return customerCount, nil
}

func (r *customerCountRepository) Delete(ctx context.Context, tx *gorm.DB, customerCountId uint64) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Delete(&entity.CustomerCount{}, "id = ?", customerCountId).Error; err != nil {
		return err
	}

	return nil
}
