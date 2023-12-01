package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/narekk1202/gin-rest-api/pkg/database"
	"github.com/narekk1202/gin-rest-api/pkg/models"
)

func GetTrackByID(ctx *gin.Context) {
	var track models.Track

	if err := database.ConnectDB().Where("id = ?", ctx.Param("id")).First(&track).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"tracks": track})
}
