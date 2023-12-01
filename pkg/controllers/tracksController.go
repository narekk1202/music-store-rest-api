package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/narekk1202/gin-rest-api/pkg/database"
	"github.com/narekk1202/gin-rest-api/pkg/models"
)

func GetAllTracks(ctx *gin.Context) {
	var tracks []models.Track

	database.ConnectDB().Find(&tracks)

	ctx.JSON(http.StatusOK, gin.H{"tracks": tracks})
}
