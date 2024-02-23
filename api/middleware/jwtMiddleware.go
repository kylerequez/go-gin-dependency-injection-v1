package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (m *Middleware) UserAuthJWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("go-gin-dependency-injection-v1-token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "Error",
				"message": "The auth cookie does not exists",
			})
			c.Abort()
			return
		}

		token, err := m.js.ValidateJWT(cookie)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "Error",
				"message": "There was an error in validating the token",
			})
			c.Abort()
			return
		}

		claims, err := m.js.ExtractJWTClaims(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "Error",
				"message": "There was an error in parsing the token",
			})
			c.Abort()
			return
		}

		exp := claims["exp"].(float64)
		if int64(exp) < time.Now().Local().Unix() {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "Error",
				"message": "The token has expired. Please login again.",
			})
			c.Abort()
			return
		}

		authority := claims["authority"].(string)
		if authority == "NORMAL_USER" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "Error",
				"message": "You are not authorized.",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
