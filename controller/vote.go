package controller

import (
	"fmt"
	"net/http"

	"github.com/emenesism/Decentralized-voting-backend/service"
	"github.com/gin-gonic/gin"
)

func GetVotes(c *gin.Context) {
	candidate := c.Query("candidate")

	if candidate == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Candidate name is required"})
		return
	}

	votes, err := service.GetVotes(candidate)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get votes"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"candidate": candidate,
		"votes":     votes.String(),
	})
}

func Vote(c *gin.Context) {
	var request struct {
		Candidate string `json:"candidate" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	tx, err := service.Vote(request.Candidate)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to submit vote"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Vote submitted successfully",
		"tx_hash": tx,
	})
}

