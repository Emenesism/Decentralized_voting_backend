package jwt_middleware


import (
	"net/http"
	"strings"
	"github.com/emenesism/Decentralized-voting-backend/models"
	"github.com/emenesism/Decentralized-voting-backend/utils/jwt"
	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearer := c.GetHeader("Authorization")
		token_str := strings.TrimPrefix(bearer, "Bearer ")

		claims, err := jwt.VerifyToken(token_str)

		if err == nil && claims != nil {
			id, ok := claims["id"].(float64)
			if !ok {
				log.Error(claims)
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				c.Abort()

				return
			}

			account := models.User{
				ID: uint(id),
			}

			if models.DB.Where(models.User{ID: account.ID}).First(&account).Error != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				c.Abort()

				return
			}

			c.Set("account", account)

		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()

			return
		}

		c.Next()
	}
}
