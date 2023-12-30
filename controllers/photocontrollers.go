
package controllers

import (
	"github.com/gin-gonic/gin"
	"PBIP_BTPN_SYARIAH/database"
	"PBIP_BTPN_SYARIAH/models"
)

func CreatePhoto(c *gin.Context) {
	var photo models.Photo
	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&photo)
	c.JSON(200, photo)
}

func GetPhotos(c *gin.Context) {
	var photos []models.Photo

	if err := database.DB.Find(&photos).Error; err != nil {
		c.AbortWithStatus(404)
	} else {
		c.JSON(200, photos)
	}
}

func UpdatePhoto(c *gin.Context) {
	photoID := c.Params.ByName("photoId")
	var photo models.Photo

	if err := database.DB.First(&photo, photoID).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}

	var newPhoto models.Photo
	if err := c.ShouldBindJSON(&newPhoto); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&photo).Updates(newPhoto)
	c.JSON(200, photo)
}

func DeletePhoto(c *gin.Context) {
	photoID := c.Params.ByName("photoId")
	var photo models.Photo

	if err := database.DB.First(&photo, photoID).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}

	database.DB.Delete(&photo)
	c.JSON(200, gin.H{"result": "Photo deleted"})
}
