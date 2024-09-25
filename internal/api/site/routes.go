package site

import (
	"github.com/gin-gonic/gin"
	"ocserv/internal/providers/middlewares"
)

func Routes(router *gin.RouterGroup) {
	site := NewSiteController()
	group := router.Group("/site")
	group.POST("/", site.Create)
	group.GET("/", site.Get)
	group.PATCH("/", middlewares.TokenMiddleware(), middlewares.IsAdminMiddleware(), site.Update)
}
