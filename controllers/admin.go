package controllers

import (
	"service-api/database"
	"service-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPendingServicer(c *gin.Context) {
	var servicers []models.Servicer
	if err := database.DB.Where("status=?", "Pending").Find(&servicers).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "no pending users",
		})
		return
	}
	c.JSON(200, gin.H{
		"pending": servicers,
	})
}

func GetAcceptedServicer(c *gin.Context) {
	var servicers []models.Servicer
	if err := database.DB.Where("status=?", "Accepted").Find(&servicers).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "no accepted users",
		})
		return
	}
	c.JSON(200, gin.H{
		"accepted": servicers,
	})
}

func GetRejectedServicer(c *gin.Context) {
	var servicers []models.Servicer
	if err := database.DB.Where("status=?", "Rejected").Find(&servicers).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "no rejected users",
		})
		return
	}
	c.JSON(200, gin.H{
		"rejected": servicers,
	})
}

func AcceptServicer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("servicer_id"))
	if err := database.DB.Model(&models.Servicer{}).Where("id=?", id).Update("status", "Accepted").Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to update status",
		})
		return
	}
	c.JSON(200, gin.H{
		"success": "Updated status",
	})
}

func RejectServicer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("servicer_id"))
	if err := database.DB.Model(&models.Servicer{}).Where("id=?", id).Update("status", "Rejected").Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to update status",
		})
		return
	}
	c.JSON(200, gin.H{
		"success": "Updated status",
	})
}
