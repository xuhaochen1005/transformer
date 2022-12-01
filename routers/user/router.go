// Package user 用户信息管理路由设置
package user

import (
	"github.com/gin-gonic/gin"
	"transformer_mz/api/user"
)

func SetupRouterGroup(rg *gin.RouterGroup) {
	userRouter := rg.Group("/user")
	// userRouter.GET("/basic", user.GetBasicUser)
	userRouter.GET("/", user.FindUsers)
	userRouter.POST("/user", user.CreateUser)
	userRouter.GET("/basic", user.GetBasicUser)
	userRouter.GET("/:id", user.GetSpecUser)
	userRouter.PATCH("/:id", user.UpdateUser)
	userRouter.DELETE("/:id", user.DeleteUser)
	userRouter.POST("/token", user.CreateUserToken)
	userRouter.PATCH("/:id/password", user.ResetUserPassword)
	userRouter.GET("/roles", user.GetUserRoles)
	userRouter.GET("/permissions", user.GetUserPermissions)
}
