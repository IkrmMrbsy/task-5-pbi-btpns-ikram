
package controllers

import (
	"github.com/gin-gonic/gin"
	"PBIP_BTPN_SYARIAH/database"
	"PBIP_BTPN_SYARIAH/models"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&user)
	c.JSON(200, user)
}

func GetUserByID(c *gin.Context) {
	var user models.User
	userID := c.Params.ByName("userId")

	if err := database.DB.Preload("Photos").First(&user, userID).Error; err != nil {
		c.AbortWithStatus(404)
	} else {
		c.JSON(200, user)
	}
}

func UpdateUser(c *gin.Context) {
	userID := c.Params.ByName("userId")
	var user models.User

	if err := database.DB.First(&user, userID).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}

	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&user).Updates(newUser)
	c.JSON(200, user)
}

func DeleteUser(c *gin.Context) {
	userID := c.Params.ByName("userId")
	var user models.User

	if err := database.DB.First(&user, userID).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}

	database.DB.Delete(&user)
	c.JSON(200, gin.H{"result": "User deleted"})
}
