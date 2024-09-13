package routes

import (
	"github.com/gin-gonic/gin"
	ApiConfigRouter "ocserv/internal/api/site"
)

// Register routers from api registering here
func Register(router *gin.RouterGroup) {
	ApiConfigRouter.Routes(router)
}
