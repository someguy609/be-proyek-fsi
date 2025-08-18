package seeds

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/someguy609/be-proyek-fsi/entity"
	"gorm.io/gorm"
)

func ListCustomerCountSeeder(db *gorm.DB) error {
	jsonFile, err := os.Open("./migrations/json/customer_count.json")
	if err != nil {
		return err
	}

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	var listCustomerCount []entity.CustomerCount
	if err := json.Unmarshal(jsonData, &listCustomerCount); err != nil {
		return err
	}

	hasTable := db.Migrator().HasTable(&entity.CustomerCount{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entity.CustomerCount{}); err != nil {
			return err
		}
	}

	for _, data := range listCustomerCount {
		var customerCount entity.CustomerCount
		err := db.Where(&entity.CustomerCount{Timestamp: data.Timestamp, LocationID: data.LocationID, Gender: data.Gender}).First(&customerCount).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		isData := db.Find(&customerCount, "timestamp = ? AND location_id = ? AND gender = ?", data.Timestamp, data.LocationID, data.Gender).RowsAffected
		if isData == 0 {
			if err := db.Create(&data).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
