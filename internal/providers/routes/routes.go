package routes

import (
	"github.com/gin-gonic/gin"
	OcservRouter "ocserv/internal/api/ocserv_user"
	ApiConfigRouter "ocserv/internal/api/site"
	UserRouter "ocserv/internal/api/user"
)

// Register routers from api registering here
func Register(router *gin.RouterGroup) {
	ApiConfigRouter.Routes(router)
	UserRouter.Routes(router)
	OcservRouter.Routes(router)
}
