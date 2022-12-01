// Package timeout 超时处理支持中间件
package timeout

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func Timeout(c *gin.Context) {
	var cancel context.CancelFunc
	beginIf, ok := c.Get("begin")
	if !ok {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError,
			"error": "内部服务错误，请联系开发人员处理"})
		c.Abort()
		return
	}
	begin := beginIf.(time.Time)
	ctx := context.Background()
	// 从header获取前端对请求处理设置的超时时间
	timeoutStr := c.Request.Header.Get("TIMEOUT")
	if timeoutStr != "" {
		timeout, err := strconv.Atoi(timeoutStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"code": http.StatusBadRequest,
				"error": "超时时间格式不正确"})
			c.Abort()
			return
		}
		deadline := begin.Add(time.Duration(timeout) * time.Second)
		if deadline.Before(begin) {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"code": http.StatusBadRequest,
				"error": "请求处理超时，请稍后再试"})
			c.Abort()
			return
		}
		ctx, cancel = context.WithDeadline(ctx, deadline)
	}
	c.Set("request-context", ctx)
	c.Set("request-cancel", cancel)
}
