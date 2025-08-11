package routes

import (
	"github.com/someguy609/be-proyek-fsi/controller"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func Location(route *gin.Engine, injector *do.Injector) {
	locationController := do.MustInvoke[controller.LocationController](injector)

	routes := route.Group("/api/location")
	{
		routes.POST("", locationController.Create)
		routes.GET("", locationController.GetAllLocation)
		routes.GET("/:id", locationController.GetLocationById)
		routes.PATCH("/:id", locationController.Update)
		routes.DELETE("/:id", locationController.Delete)
	}
}
