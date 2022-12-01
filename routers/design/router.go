// Package design 设计任务管理
package design

import (
	"transformer_mz/api/approve"
	"transformer_mz/api/check"
	"transformer_mz/api/design"
	"transformer_mz/api/review"
	"transformer_mz/api/submit"

	"github.com/gin-gonic/gin"
)

func SetupRouterGroup(rg *gin.RouterGroup) {
	designRouter := rg.Group("/design")

	designRouter.GET("/projects", design.FindDesignProjects)
	designRouter.POST("/project", design.CreateDesignProjects)
	designRouter.GET("/project/:id", design.GetSpecDesignProjects)
	designRouter.PATCH("/project/:id", design.UpdateSpecDesignProjects)
	designRouter.DELETE("/project/:id", design.DeleteSpecDesignProjects)

	designRouter.POST("/project_export/:id", design.ProjectsExport)

	designRouter.GET("/tasks", design.FindDesignTasks)
	designRouter.GET("/tasks/results", design.FindTaskDesignResults)
	designRouter.POST("/task", design.CreateDesignTask)
	designRouter.GET("/task/:id", design.GetSpecDesignTask)
	designRouter.PATCH("/task_submit", submit.UpdateDesignTaskStatus)
	designRouter.PATCH("/task_review", review.UpdateDesignTaskStatus)
	designRouter.PATCH("/task_approve", approve.UpdateDesignTaskStatus)
	designRouter.PATCH("/task_check", check.UpdateDesignTaskStatus)
	designRouter.PATCH("/task/:id", design.UpdateDesignTask)
	designRouter.GET("/task/export/:id", design.TasksExcelExport)

	designRouter.GET("/start", design.StartAIDesign)
	designRouter.POST("/cancel", design.CanceAIDesign)
	designRouter.GET("/hook/:id", design.CreateAIDesignHook)

	designRouter.GET("/recommand", design.RecommendDesignTasks)

	designRouter.GET("/design_results", design.FindDesignResults)
	designRouter.POST("/design_result", design.CreateDesignResult)
	designRouter.POST("/design_results", design.CreateDesignResults)
	designRouter.GET("/design_result/:id", design.GetSpecDesignResults)
	designRouter.PATCH("/design_result/:id", design.UpdateSpecDesignResults)
	designRouter.DELETE("/design_result/:id", design.DeleteSpecDesignResults)
	designRouter.POST("/design_result/:id", design.DesignResultsExport)
}
