package seeds

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/someguy609/be-proyek-fsi/entity"
	"gorm.io/gorm"
)

func ListLocationSeeder(db *gorm.DB) error {
	jsonFile, err := os.Open("./migrations/json/locations.json")
	if err != nil {
		return err
	}

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	var listLocation []entity.Location
	if err := json.Unmarshal(jsonData, &listLocation); err != nil {
		return err
	}

	hasTable := db.Migrator().HasTable(&entity.Location{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entity.Location{}); err != nil {
			return err
		}
	}

	for _, data := range listLocation {
		var location entity.Location
		err := db.Where(&entity.Location{Name: data.Name}).First(&location).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		isData := db.Find(&location, "name = ?", data.Name).RowsAffected
		if isData == 0 {
			if err := db.Create(&data).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
