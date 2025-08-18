package routes

import (
	"github.com/someguy609/be-proyek-fsi/controller"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func CustomerCount(route *gin.Engine, injector *do.Injector) {
	customerCountController := do.MustInvoke[controller.CustomerCountController](injector)

	routes := route.Group("/api/location")
	{
		routes.POST("/:id/customer-count", customerCountController.Create)
		routes.GET("/:id/customer-count", customerCountController.GetCustomerCountByLocation)
		routes.PATCH("/:id/customer-count", customerCountController.Update)
		// routes.DELETE("/:id/customer-count", customerCountController.Delete)
	}
}
