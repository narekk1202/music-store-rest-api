package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/narekk1202/gin-rest-api/pkg/database"
	"github.com/narekk1202/gin-rest-api/pkg/models"
)

func UpdateTrack(ctx *gin.Context) {
	var track models.Track
	DB := database.ConnectDB()

	if err := DB.Where("id = ?", ctx.Param("id")).First(&track).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}

	var updatedTrack models.UpdateTrack

	if err := ctx.ShouldBindJSON(&updatedTrack); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body"})
		return
	}

	DB.Model(&track).Updates(updatedTrack)
	ctx.JSON(http.StatusOK, gin.H{"tracks": track})
}
