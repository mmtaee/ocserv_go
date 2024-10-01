package routes

import (
	"github.com/gin-gonic/gin"
	OcservGroupRouter "ocserv/internal/api/ocserv_group"
	OcservUserRouter "ocserv/internal/api/ocserv_user"
	ApiConfigRouter "ocserv/internal/api/site"
	UserRouter "ocserv/internal/api/user"
)

// Register routers from api registering here
func Register(router *gin.RouterGroup) {
	ApiConfigRouter.Routes(router)
	UserRouter.Routes(router)
	OcservUserRouter.Routes(router)
	OcservGroupRouter.Routes(router)
}
