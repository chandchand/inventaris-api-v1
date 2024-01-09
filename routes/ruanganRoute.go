package routes

import (
	"example/inventaris-api-v1/controllers"
	"example/inventaris-api-v1/services"

	"github.com/gin-gonic/gin"
)

func SetupRuangRoute(r *gin.Engine, ruangService *services.RuangService) {
	ruangController := controllers.NewRuangController(ruangService)

	routeGroup := r.Group("/api/v1")
	{
		routeGroup.POST("/ruangan", ruangController.Create)
		routeGroup.GET("/ruangan", ruangController.GetAll)
		routeGroup.GET("/ruangan/:id", ruangController.GetOne)
		routeGroup.PUT("/ruangan/:id", ruangController.Update)
		routeGroup.DELETE("/ruangan/:id", ruangController.Delete)
	}
}