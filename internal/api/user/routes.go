package user

import "github.com/gin-gonic/gin"

func Routes(router *gin.RouterGroup) {
	user := NewUserController()
	group := router.Group("/user")
	group.POST("/", user.CreateAdminUser)
	group.PATCH("/password/", user.UpdatePassword)
}
