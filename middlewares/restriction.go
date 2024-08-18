package middlewares

import (
	database "final-project/config/postgres"
	"final-project/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func RestrictMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()

		var profileCount int64

		// Data User Define
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		db.Model(&models.Profile{}).Where("user_id = ?", userID).Count(&profileCount)

		if profileCount == 0 {
			c.JSON(http.StatusForbidden, gin.H{
				"error":   "Forbidden",
				"message": "Complete your profile first.",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
