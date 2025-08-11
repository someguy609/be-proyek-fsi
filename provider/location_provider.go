package provider

import (
	"github.com/someguy609/be-proyek-fsi/controller"
	"github.com/someguy609/be-proyek-fsi/repository"
	"github.com/someguy609/be-proyek-fsi/service"
	"github.com/samber/do"
	"gorm.io/gorm"
)

func ProvideLocationDependencies(injector *do.Injector, db *gorm.DB) {
	// Repository
	locationRepository := repository.NewLocationRepository(db)

	// Service
	locationService := service.NewLocationService(locationRepository, db)

	// Controller
	do.Provide(
		injector, func(i *do.Injector) (controller.LocationController, error) {
			return controller.NewLocationController(locationService), nil
		},
	)
}