package main

import (
	"github.com/gin-gonic/gin"
	"github.com/narekk1202/music-store-rest-api/pkg/database"
	"github.com/narekk1202/music-store-rest-api/pkg/routes"
)

func main() {
	router := gin.Default()

	database.ConnectDB()
	routes.SetupRoutes(router)

	router.Run()
}
