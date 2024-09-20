package user

import (
	"github.com/gin-gonic/gin"
	"ocserv/internal/providers/routes/middleware"
)

func Routes(router *gin.RouterGroup) {
	user := NewUserController()
	group := router.Group("/users")
	group.POST("/", user.CreateAdminUser)
	group.POST("/login/", user.Login)
	group.POST("/password/", middleware.TokenMiddleware(), user.UpdatePassword)

	staffGroup := group.Group("/staff", middleware.TokenMiddleware())
	staffGroup.POST("/staffs/", user.CreateStaff)
	staffGroup.POST("/staffs/:id/password/", user.UpdateStaffPassword)
	staffGroup.DELETE("staffs/:id/", user.DeleteStaff)
}
