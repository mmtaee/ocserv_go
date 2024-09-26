package ocserv

import (
	"github.com/gin-gonic/gin"
	"ocserv/internal/providers/middlewares"
)

func Routes(router *gin.RouterGroup) {
	ocserv := NewOcservController()
	ocservGroup := router.Group("/ocserv")
	ocservGroup.Use(middlewares.TokenMiddleware())
	ocservGroup.POST("/", ocserv.Create)
	ocservGroup.PATCH("/:id/", ocserv.Update)
	ocservGroup.DELETE("/:id/", ocserv.Delete)
	ocservGroup.POST("/:id/disconnect/", ocserv.Disconnect)
}
