package ocserv_group

import "github.com/gin-gonic/gin"

func Routes(router *gin.RouterGroup) {
	ocservGroup := NewOcservGroupController()
	group := router.Group("/ocserv/groups")
	group.GET("/", ocservGroup.List)
	group.POST("/", ocservGroup.Create)
	group.PATCH("/:name/", ocservGroup.Update)
	group.DELETE("/:name/", ocservGroup.Delete)
}
