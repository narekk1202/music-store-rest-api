package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/narekk1202/gin-rest-api/pkg/database"
	"github.com/narekk1202/gin-rest-api/pkg/models"
)

func CreateTrack(ctx *gin.Context) {
	var createdTrack models.CreateTrack
	DB := database.ConnectDB() 

	if err := ctx.ShouldBindJSON(&createdTrack); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body"})
		return
	}

	track := models.Track{Artist: createdTrack.Artist, Title: createdTrack.Title}
	DB.Create(&track)
	ctx.JSON(http.StatusOK, gin.H{"tracks": track})
}
