package controller

import (
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/emenesism/Decentralized-voting-backend/models"
	"github.com/emenesism/Decentralized-voting-backend/utils/security"
	"github.com/gin-gonic/gin"
	"github.com/emenesism/Decentralized-voting-backend/utils/jwt"
)

func Register(c *gin.Context) {
	var request struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Invalid request data", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	hashedPassword := security.MD5Hash(request.Password)

	user := models.User{
		Username: request.Username,
		Hash:     hashedPassword,
	}

	if err := models.DB.Create(&user).Error; err != nil {
		log.Error("Failed to create user", "username", request.Username, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	log.Info("User registered successfully", "username", request.Username)
	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
		},
	})
}

func Login(c *gin.Context) {
	var request struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Invalid request data", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	var user models.User

	if err := models.DB.Where("username = ?", request.Username).First(&user).Error; err != nil {
		log.Error("User not found", "username", request.Username, "error", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	hashedPassword := security.MD5Hash(request.Password)

	if user.Hash != hashedPassword {
		log.Error("Invalid password", "username", request.Username)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	log.Info("User logged in successfully", "username", request.Username)

	token , err := jwt.GenToken(user.ID)

	if err != nil {
		log.Error("Failed to generate token", "username", request.Username, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"token": token,
		},
	})

}
