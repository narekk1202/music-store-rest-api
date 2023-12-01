package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/narekk1202/music-store-rest-api/pkg/controllers"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/tracks", controllers.GetAllTracks)
	r.GET("/tracks/:id", controllers.GetTrackByID)
	r.POST("/create", controllers.CreateTrack)
	r.DELETE("/delete/:id", controllers.DeleteTrack)
	r.PATCH("/update/:id", controllers.UpdateTrack)
}