// Package department 文档管理路由设置
package documents

import (
	"github.com/gin-gonic/gin"
	Documents "transformer_mz/api/documents"
)

func SetupRouterDocuments(rg *gin.RouterGroup) {
	documentsRouter := rg.Group("/documents")
	documentsRouter.GET("/", Documents.FindDocuments)
	documentsRouter.POST("/upload", Documents.UploadDocuments)
	documentsRouter.GET("/download/:id", Documents.DownloadDocuments)
	documentsRouter.GET("/:id", Documents.GetSpecDocuments)
	documentsRouter.PATCH("/:id", Documents.UpdateDocuments)
	documentsRouter.DELETE("/:id", Documents.DeleteDocuments)
}
