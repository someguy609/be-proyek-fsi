package repository

// import (
// 	"context"
// 	"time"

// 	"github.com/someguy609/be-proyek-fsi/entity"
// 	"gorm.io/gorm"
// )

// type (
// 	CustomerRepository interface {
// 		Create(ctx context.Context, tx *gorm.DB, customer entity.Customer, location entity.Location) (entity.Customer, error)
// 		GetCustomerById(ctx context.Context, tx *gorm.DB, customerId string) (entity.Customer, error)
// 		GetCustomerCountByLocation(ctx context.Context, tx *gorm.DB, locationId string, start *time.Time, end *time.Time, interval string) ([]entity.Customer, error)
// 		Update(ctx context.Context, tx *gorm.DB, customerId entity.Customer) (entity.Customer, error)
// 		Delete(ctx context.Context, tx *gorm.DB, customerId string) error
// 	}

// 	customerRepository struct {
// 		db *gorm.DB
// 	}
// )

// func NewCustomerRepository(db *gorm.DB) CustomerRepository {
// 	return &customerRepository{
// 		db: db,
// 	}
// }

// func (r *customerRepository) Create(ctx context.Context, tx *gorm.DB, customer entity.Customer, location entity.Location) (entity.Customer, error) {
// 	if tx == nil {
// 		tx = r.db
// 	}

// 	if err := tx.WithContext(ctx).Create(&customer).Error; err != nil {
// 		return entity.Customer{}, err
// 	}

// 	if err := tx.WithContext(ctx).Association("locations").Append(&location); err != nil {
// 		return entity.Customer{}, err
// 	}

// 	return customer, nil
// }

// func (r *customerRepository) GetCustomerById(ctx context.Context, tx *gorm.DB, customerId string) (entity.Customer, error) {
// 	if tx == nil {
// 		tx = r.db
// 	}

// 	if err := tx.WithContext(ctx).Model(&entity.Customer{}).Where("id = ?", customerId); err != nil {
// 		return entity.Customer{}, nil
// 	}

// 	return entity.Customer{}, nil
// }

// func (r *customerRepository) GetCustomerCountByLocation(
// 	ctx context.Context,
// 	tx *gorm.DB,
// 	locationId string,
// 	start *time.Time,
// 	end *time.Time,
// 	interval string,
// ) ([]entity.Customer, error) {
// 	if tx == nil {
// 		tx = r.db
// 	}

// 	query := tx.Raw(`
// 		SELECT
// 			time_bucket(?, timestamp) AS timestamp,
// 			gender,
// 			COUNT(DISTINCT customer_id) as count
// 		FROM
// 			customer_counts
// 		WHERE
// 			location_id = ? AND timestamp >= ? AND timestamp <= ?
// 		GROUP BY
// 			timestamp, gender
// 		ORDER BY
// 			timestamp
// 	`, interval, locationId, start, end)

// 	var customers []entity.Customer

// 	if err := query.Scan(&customers).Error; err != nil {
// 		return nil, err
// 	}

// 	return customers, nil
// }

// func (r *customerRepository) Update(ctx context.Context, tx *gorm.DB, customer entity.Customer) (entity.Customer, error) {
// 	if tx == nil {
// 		tx = r.db
// 	}

// 	if err := tx.WithContext(ctx).Updates(&customer).Error; err != nil {
// 		return entity.Customer{}, err
// 	}

// 	return customer, nil
// }

// func (r *customerRepository) Delete(ctx context.Context, tx *gorm.DB, customerId string) error {
// 	if tx == nil {
// 		tx = r.db
// 	}

// 	if err := tx.WithContext(ctx).Delete(&entity.Customer{}, "id = ?", customerId).Error; err != nil {
// 		return err
// 	}

// 	return nil
// }
