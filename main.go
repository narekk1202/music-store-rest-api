package main

import (
	"github.com/gin-gonic/gin"
	"github.com/narekk1202/gin-rest-api/pkg/controllers"
	"github.com/narekk1202/gin-rest-api/pkg/database"
)

func main() {
	r := gin.Default()

	database.ConnectDB()

	r.GET("/tracks", controllers.GetAllTracks)
	r.GET("/tracks/:id", controllers.GetTrackByID)
	r.POST("/create", controllers.CreateTrack)
	r.DELETE("/delete/:id", controllers.DeleteTrack)
	r.PATCH("/update/:id", controllers.UpdateTrack)
	r.Run()
}
