package controllers

import (
	"net/http"
	"test_case/database"
	"test_case/models"

	"github.com/gin-gonic/gin"
)

func GetUserData(c *gin.Context) {
	iduser := c.Param("iduser")

	var user models.User
	if err := database.DB.Preload("Addresses").Where("iduser = ?", iduser).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func UpdateUserAddress(c *gin.Context) {
	iduser := c.Param("iduser")
	var input models.Address

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx := database.DB.Begin()

	if err := tx.Model(&models.Address{}).Where("iduser = ?", iduser).Updates(input).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"message": "Address updated successfully"})
}
