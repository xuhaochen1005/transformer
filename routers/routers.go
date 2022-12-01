// Package routers 全局路由设置
package routers

import (
	"github.com/gin-gonic/gin"
	"transformer_mz/middleware/auth"
	"transformer_mz/middleware/timeout"
	"transformer_mz/middleware/trace"
	"transformer_mz/routers/approve"
	"transformer_mz/routers/department"
	"transformer_mz/routers/design"
	"transformer_mz/routers/documents"
	"transformer_mz/routers/feedback"
	"transformer_mz/routers/maintenance"
	"transformer_mz/routers/message"
	"transformer_mz/routers/permission"
	"transformer_mz/routers/standard"
	"transformer_mz/routers/statistics"
	"transformer_mz/routers/todolist"
	"transformer_mz/routers/user"
)

func SetupRouter() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)
	router := gin.New()
	// 设置全局中间件
	router.Use(auth.JWTAuth)
	router.Use(timeout.Timeout)
	router.Use(trace.Trace)
	v1 := router.Group("/v1")
	approve.SetupRouterGroup(v1)
	department.SetupRouterGroup(v1)
	design.SetupRouterGroup(v1)
	message.SetupRouterGroup(v1)
	permission.SetupRouterGroup(v1)
	standard.SetupRouterGroup(v1)
	user.SetupRouterGroup(v1)
	maintenance.SetupRouterGroup(v1)
	documents.SetupRouterDocuments(v1)
	statistics.SetupStatisticsGroup(v1)
	todolist.SetupTodoListGroup(v1)
	feedback.SetupRouterGroup(v1)
	return router
}
