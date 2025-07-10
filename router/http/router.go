package http

import (
	"github.com/emenesism/Decentralized-voting-backend/middleware/logger"
	"github.com/emenesism/Decentralized-voting-backend/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(logger.CustomLogger())

	v1 := router.Group("v1")
	{
		v1.GET("health", controller.HealthCheck)
		v1.GET("votes", controller.GetVotes)
		v1.POST("vote", controller.Vote)	
	}

	return router
}
