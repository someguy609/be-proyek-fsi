package provider

import (
	"github.com/someguy609/be-proyek-fsi/controller"
	"github.com/someguy609/be-proyek-fsi/repository"
	"github.com/someguy609/be-proyek-fsi/service"
	"github.com/samber/do"
	"gorm.io/gorm"
)

func ProvideCustomerCountDependencies(injector *do.Injector, db *gorm.DB) {
	// Repository
	customerCountRepository := repository.NewCustomerCountRepository(db)

	// Service
	customerCountService := service.NewCustomerCountService(customerCountRepository, db)

	// Controller
	do.Provide(
		injector, func(i *do.Injector) (controller.CustomerCountController, error) {
			return controller.NewCustomerCountController(customerCountService), nil
		},
	)
}