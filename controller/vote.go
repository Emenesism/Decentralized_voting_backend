package controller

import (
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/emenesism/Decentralized-voting-backend/service"
	"github.com/gin-gonic/gin"
)

func GetVotes(c *gin.Context) {
	candidate := c.Query("candidate")

	if candidate == "" {
		log.Error("Candidate name is missing in the request")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Candidate name is required"})
		return
	}

	votes, err := service.GetVotes(candidate)
	if err != nil {
		log.Error("Failed to get votes", "candidate", candidate, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get votes"})
		return
	}

	log.Info("Successfully retrieved votes", "candidate", candidate, "votes", votes.String())
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
		log.Error("Invalid request data", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	tx, err := service.Vote(request.Candidate)
	if err != nil {
		log.Error("Failed to submit vote", "candidate", request.Candidate, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to submit vote"})
		return
	}

	log.Info("Vote submitted successfully", "candidate", request.Candidate, "tx_hash", tx)
	c.JSON(http.StatusOK, gin.H{
		"message": "Vote submitted successfully",
		"tx_hash": tx,
	})
}
