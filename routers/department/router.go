// Package department 部门信息管理路由设置
package department

import (
	"github.com/gin-gonic/gin"
	"transformer_mz/api/department"
)

func SetupRouterGroup(rg *gin.RouterGroup) {
	departmentRouter := rg.Group("/department")
	departmentRouter.GET("/", department.FindDepartments)
	departmentRouter.POST("/", department.CreateDepartment)
	departmentRouter.GET("/:id", department.GetSpecDepartment)
	departmentRouter.PATCH("/:id", department.UpdateDepartment)
	departmentRouter.DELETE("/:id", department.DeleteDepartment)
	departmentRouter.POST("/export", department.ExcelExport)
}
