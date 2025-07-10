package http

import (
	"github.com/emenesism/Decentralized-voting-backend/controller"
	jwt_middleware "github.com/emenesism/Decentralized-voting-backend/middleware/jwt"
	"github.com/emenesism/Decentralized-voting-backend/middleware/logger"
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
		v1.POST("vote", jwt_middleware.AuthMiddleware(), controller.Vote)
		v1.POST("register", controller.Register)
		v1.POST("login", controller.Login)
	}

	return router
}
