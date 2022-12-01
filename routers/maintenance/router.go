// Package 标准库 用户信息管理路由设置
package maintenance

import (
	"github.com/gin-gonic/gin"
	"transformer_mz/api/maintenance"
)

func SetupRouterGroup(rg *gin.RouterGroup) {
	maintenanceRouter := rg.Group("/sys")
	// 系统监控维护接口
	maintenanceRouter.GET("/host", maintenance.SystemHostsInfoStatus)
	maintenanceRouter.GET("/basic", maintenance.SystemBasicInfoStatus)
	maintenanceRouter.GET("/cpu", maintenance.SystemCpuInfoInfoStatus)
	maintenanceRouter.GET("/disk", maintenance.SystemDisksInfoStatus)
	maintenanceRouter.GET("/net", maintenance.SystemNetInfoStatus)
	maintenanceRouter.GET("/process", maintenance.SystemProcessInfoStatus)
	maintenanceRouter.GET("/mem", maintenance.SystemMemInfoStatus)
}
