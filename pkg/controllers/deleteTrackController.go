package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/narekk1202/music-store-rest-api/pkg/database"
	"github.com/narekk1202/music-store-rest-api/pkg/models"
)

func DeleteTrack(ctx *gin.Context) {
	var track models.Track
	DB := database.ConnectDB()

	if err := DB.Where("id = ?", ctx.Param("id")).First(&track).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}

	DB.Delete(&track)
	ctx.JSON(http.StatusOK, gin.H{"message": "Запись успешно удалена"})
}
