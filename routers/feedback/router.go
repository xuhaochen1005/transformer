package feedback

import (
	"github.com/gin-gonic/gin"
	"transformer_mz/api/feedback"
)

func SetupRouterGroup(rg *gin.RouterGroup) {
	feedbackGroup := rg.Group("/feedback")
	feedbackGroup.GET("/", feedback.FindFeedback)
	feedbackGroup.GET("/:id", feedback.GetSpecFeedback)
	feedbackGroup.POST("/", feedback.CreateFeedback)
	feedbackGroup.DELETE("/:id", feedback.DeleteFeedback)
}
