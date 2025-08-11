package provider

import (
	"github.com/someguy609/be-proyek-fsi/controller"
	"github.com/someguy609/be-proyek-fsi/repository"
	"github.com/someguy609/be-proyek-fsi/service"
	"github.com/samber/do"
	"gorm.io/gorm"
)

func ProvideUserDependencies(injector *do.Injector, db *gorm.DB, jwtService service.JWTService) {
	// Repository
	userRepository := repository.NewUserRepository(db)
	refreshTokenRepository := repository.NewRefreshTokenRepository(db)

	// Service
	userService := service.NewUserService(userRepository, refreshTokenRepository, jwtService, db)

	// Controller
	do.Provide(
		injector, func(i *do.Injector) (controller.UserController, error) {
			return controller.NewUserController(userService), nil
		},
	)
}
