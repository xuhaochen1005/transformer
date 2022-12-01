package statistics

import (
	"github.com/gin-gonic/gin"
	"transformer_mz/api/statistics"
)

func SetupStatisticsGroup(rg *gin.RouterGroup) {
	statisticsRouter := rg.Group("/statistics")

	//dashboard统计数据
	statisticsRouter.GET("/", statistics.GetStatistics)
}
