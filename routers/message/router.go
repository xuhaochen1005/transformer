// Package notify 消息通知管理
package message

import (
	"github.com/gin-gonic/gin"
	"transformer_mz/api/message"
)

func SetupRouterGroup(rg *gin.RouterGroup) {
	messageRouter := rg.Group("/notify")
	messageRouter.GET("/message", message.FindMessages)
	messageRouter.POST("/notify", message.Notify)
	messageRouter.POST("/message/UpdateAllMessageRead", message.UpdateAllMessageRead)
	messageRouter.GET("/message/:id", message.GetSpecMessage)
	messageRouter.PATCH("/message/:id", message.UpdateMessageStatus)
	messageRouter.DELETE("/message/:id", message.DeleteMessage)
	messageRouter.GET("/message_template", message.FindMessageTemplates)
	messageRouter.POST("/message_template", message.CreateMessageTemplate)
	messageRouter.GET("/message_template/:id", message.GetSpecMessageTemplate)
	messageRouter.PATCH("/message_template/:id", message.UpdateMessageTemplate)
	messageRouter.DELETE("/message_template/:id", message.DeleteMessageTemplate)
}
