package ocserv_group

import (
	"github.com/gin-gonic/gin"
	"ocserv/internal/providers/middlewares"
)

func Routes(router *gin.RouterGroup) {
	ocservGroup := NewOcservGroupController()
	group := router.Group("/ocserv/groups")
	group.Use(middlewares.TokenMiddleware())
	group.GET("/", ocservGroup.List)
	group.POST("/", ocservGroup.Create)
	group.PATCH("/:name/", ocservGroup.Update)
	group.DELETE("/:name/", ocservGroup.Delete)
}
