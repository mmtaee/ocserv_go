package ocserv_group

import "github.com/gin-gonic/gin"

func Routes(router *gin.RouterGroup) {
	ocservGroup := NewOcservGroupController()
	group := router.Group("/ocserv/groups")
	group.GET("/", ocservGroup.List)
	group.POST("/", ocservGroup.Create)
	group.PATCH("/:id/", ocservGroup.Update)
	group.DELETE("/:id/", ocservGroup.Delete)
}
