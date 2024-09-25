package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
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
		tokenKey = strings.TrimPrefix(tokenKey, "Bearer ")
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
		log.Println(user)
		c.Set("userId", user.ID)
		c.Set("user", user)
		c.Set("isAdmin", user.IsAdmin)
		c.Next()
	}
}
