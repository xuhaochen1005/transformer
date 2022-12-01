// Package permission 权限管理
package permission

import (
	"github.com/gin-gonic/gin"
	"transformer_mz/api/permission"
)

func SetupRouterGroup(rg *gin.RouterGroup) {
	PermissionRouter := rg.Group("/permission")
	PermissionRouter.GET("/role_permission", permission.FindRolePermission)
	PermissionRouter.POST("/role_permission", permission.CreateRolePermission)
	PermissionRouter.GET("/role_permission/:id", permission.GetSpecRolePermission)
	PermissionRouter.DELETE("/role_permission/:id", permission.DeleteRolePermission)
	PermissionRouter.GET("/user_role", permission.FindUserRole)
	PermissionRouter.POST("/user_role", permission.CreateUserRole)
	PermissionRouter.GET("/user_role/:id", permission.GetSpecUserRole)
	PermissionRouter.PATCH("/user_role/:id", permission.UpdateUserRole)
	PermissionRouter.DELETE("/user_role/:id", permission.DeleteUserRole)
	PermissionRouter.GET("/roles", permission.FindRoles)
	PermissionRouter.POST("/role", permission.CreateRole)
	PermissionRouter.PATCH("/role/:id", permission.UpdateRole)
	PermissionRouter.DELETE("/role/:id", permission.DeleteRole)

}
