package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ocserv/internal/repository"
	"strings"
	"time"
)

func unauthorized(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"error": "Invalid Authorization Credentials",
	})
}

func TokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenKey := c.Request.Header.Get("Authorization")
		tokenKey = strings.TrimPrefix(tokenKey, "Token ")
		if tokenKey == "" {
			unauthorized(c)
			return
		}
		tokenRepository := repository.NewTokenRepository()
		user, token, err := tokenRepository.GetTokenByKey(tokenKey)
		if err != nil {
			unauthorized(c)
			return
		}
		if token.ExpireAt <= time.Now().Unix() {
			unauthorized(c)
			return
		}
		c.Set("userId", user.ID)
		c.Set("isStaff", user.IsStaff)
		c.Next()
	}
}
