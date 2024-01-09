package main

import (
	"example/inventaris-api-v1/config"
	"example/inventaris-api-v1/repositories"
	"example/inventaris-api-v1/routes"
	"example/inventaris-api-v1/services"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	db, err := cfg.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	ruangRepo := repositories.NewRuangRepo(db)
	InventarisRepo := repositories.NewInventRepo(db)

	ruangService := services.NewRuangService(ruangRepo)
	inventarisService := services.NewInventService(InventarisRepo)

	r := gin.Default()

	routes.SetupRuangRoute(r, ruangService)
	routes.SetupInventRoute(r, inventarisService)

	r.Run(":8080")


}