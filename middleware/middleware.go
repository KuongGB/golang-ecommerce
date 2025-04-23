package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/tokens"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		ClientToken := c.Request.Header.Get("token")
		if ClientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "token required"})
			c.Abort()
			return
		}
		claims, err := tokens.ValidateToken(ClientToken)
		if err != "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err})
			c.Abort()
			return
		}
		c.Set("uid", claims.Uid)
		c.Set("email", claims.Email)
		c.Next()
	}
}
