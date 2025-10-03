package provider

import (
	"github.com/samber/do"
	"github.com/someguy609/be-proyek-fsi/controller"
	"github.com/someguy609/be-proyek-fsi/repository"
	"github.com/someguy609/be-proyek-fsi/service"
	"go.mongodb.org/mongo-driver/v2/mongo"
	// "gorm.io/gorm"
)

func ProvideCustomerCountDependencies(injector *do.Injector, db *mongo.Database) {
	// Repository
	locationRepository := repository.NewLocationRepository(db)
	customerCountRepository := repository.NewCustomerCountRepository(db)

	// Service
	customerCountService := service.NewCustomerCountService(customerCountRepository, locationRepository, db)

	// Controller
	do.Provide(
		injector, func(i *do.Injector) (controller.CustomerCountController, error) {
			return controller.NewCustomerCountController(customerCountService), nil
		},
	)
}