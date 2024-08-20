package controllers

import (
	"net/http"
	"test_case/auth"
	"test_case/database"
	"test_case/models"

	"github.com/gin-gonic/gin"
)

func GetUserData(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	const BEARER_SCHEMA = "Bearer "
	tokenString = tokenString[len(BEARER_SCHEMA):]
	parse, err := auth.ParseToken(tokenString)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not login"})
		return
	}

	var user models.User
	if err := database.DB.Preload("Addresses").Where("id_user = ?", parse.IdUser).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func UpdateUserAddress(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	const BEARER_SCHEMA = "Bearer "
	tokenString = tokenString[len(BEARER_SCHEMA):]
	parse, err := auth.ParseToken(tokenString)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not login"})
		return
	}

	var input models.Address

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx := database.DB.Begin()

	if err := tx.Model(&models.Address{}).Where("id_user = ?", parse.IdUser).Updates(input).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"message": "Address updated successfully"})
}
