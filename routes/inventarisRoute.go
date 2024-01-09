package routes

import (
	"example/inventaris-api-v1/controllers"
	"example/inventaris-api-v1/services"

	"github.com/gin-gonic/gin"
)

func SetupInventRoute(r *gin.Engine, inventarisService *services.InventarisService) {
	
	InventarisController := controllers.NewInventController(inventarisService)

	routeGroup := r.Group("/api/v1")
	{
		routeGroup.POST("/inventaris", InventarisController.Create)
		routeGroup.GET("/inventaris", InventarisController.GetAll)
		routeGroup.GET("/inventaris/search", InventarisController.Search)
		routeGroup.GET("/inventaris/:id", InventarisController.GetOne)
		routeGroup.PUT("/inventaris/:id", InventarisController.Update)
		routeGroup.DELETE("/inventaris/:id", InventarisController.Delete)
	}
}