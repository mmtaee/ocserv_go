package user

import (
	"github.com/gin-gonic/gin"
	"ocserv/internal/providers/middlewares"
)

func Routes(router *gin.RouterGroup) {
	user := NewUserController()
	group := router.Group("/users")
	group.POST("/", user.CreateAdminUser)
	group.POST("/login/", user.Login)
	group.PATCH("/password/", middlewares.TokenMiddleware(), user.UpdatePassword)

	staffGroup := group.Group("/", middlewares.TokenMiddleware())
	staffGroup.POST("/staffs/", middlewares.IsAdminMiddleware(), user.CreateStaff)
	staffGroup.PATCH("/staffs/:id/password/", middlewares.IsAdminMiddleware(), user.UpdateStaffPassword)
	staffGroup.DELETE("staffs/:id/", middlewares.IsAdminMiddleware(), user.DeleteStaff)
}
