package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IsAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		isAdmin, _ := c.Get("isAdmin")
		if !isAdmin.(bool) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Admin Permission required"})
			return
		}
		c.Next()
	}
}
