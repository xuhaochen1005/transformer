// Package approve 审批管理
package approve

import (
	"github.com/gin-gonic/gin"
	"transformer_mz/api/approve"
)

func SetupRouterGroup(rg *gin.RouterGroup) {
	approveRouter := rg.Group("/approve")

	approveRouter.GET("/approves", approve.FindApproves)
	approveRouter.POST("/approve", approve.CreateApprove)
	approveRouter.GET("/approve/:id", approve.GetSpecApprove)
	approveRouter.PATCH("/approve/:id", approve.UpdateApprove)
	approveRouter.DELETE("/approve/:id", approve.DeleteApprove)

	approveRouter.GET("/approve_nodes", approve.FindApproveNodes)
	approveRouter.POST("/approve_node", approve.CreateApproveNode)
	approveRouter.GET("/approve_node/:id", approve.GetSpecApproveNode)
	approveRouter.PATCH("/approve_node/:id", approve.UpdateApproveNode)
	approveRouter.DELETE("/approve_node/:id", approve.DeleteApproveNode)

	approveRouter.GET("/approve_instances", approve.FindApproveInstances)
	approveRouter.POST("/approve_instance", approve.CreateApproveInstance)
	approveRouter.GET("/approve_instance/:id", approve.GetSpecApproveInstance)

	approveRouter.GET("/approve_node_instances", approve.FindApproveNodeInstances)
	approveRouter.GET("/approve_node_instance/:id", approve.GetSpecApproveNodeInstance)
	approveRouter.PATCH("/approve_node_instance/:id", approve.UpdateApproveNodeInstance)
}
