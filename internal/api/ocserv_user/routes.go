package ocserv_user

import (
	"github.com/gin-gonic/gin"
	"ocserv/internal/providers/middlewares"
)

func Routes(router *gin.RouterGroup) {
	ocservUser := NewOcservUserController()
	group := router.Group("/ocserv/users/")
	group.Use(middlewares.TokenMiddleware())
	group.POST("/", ocservUser.Create)
	group.PATCH("/:id/", ocservUser.Update)
	group.DELETE("/:id/", ocservUser.Delete)
	group.POST("/:id/disconnect/", ocservUser.Disconnect)
}
