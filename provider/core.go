package provider

import (
	"github.com/samber/do"
	"github.com/someguy609/be-proyek-fsi/config"
	"github.com/someguy609/be-proyek-fsi/constants"
	"github.com/someguy609/be-proyek-fsi/service"
	"go.mongodb.org/mongo-driver/v2/mongo"
	// "gorm.io/gorm"
)

func InitDatabase(injector *do.Injector) {
	do.ProvideNamed(injector, constants.DB, func(i *do.Injector) (*mongo.Database, error) {
		return config.SetUpDatabaseConnection(), nil
	})
}

func RegisterDependencies(injector *do.Injector) {
	InitDatabase(injector)

	do.ProvideNamed(injector, constants.JWTService, func(i *do.Injector) (service.JWTService, error) {
		return service.NewJWTService(), nil
	})

	// Initialize
	db := do.MustInvokeNamed[*mongo.Database](injector, constants.DB)
	// jwtService := do.MustInvokeNamed[service.JWTService](injector, constants.JWTService)

	// Provide Dependencies
	// ProvideUserDependencies(injector, db, jwtService)
	ProvideLocationDependencies(injector, db)
	ProvideCustomerCountDependencies(injector, db)
}
