// Package trace 请求信息跟踪中间件
package trace

import (
	"context"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	"transformer_mz/libs/errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"transformer_mz/databases/model"
	"transformer_mz/libs/log"
)

func Trace(c *gin.Context) {
	var requestID string
	// 生成请求的唯一标识
	id, err := uuid.NewUUID()
	if err == nil {
		requestID = id.String()
	} else {
		requestID = strconv.Itoa(rand.New(rand.NewSource(time.Now().UnixNano())).Int())
	}
	// 将请求的唯一标识设置到返回的header中
	c.Header("X-Request-Id", requestID)
	username := "???"
	user, ok := c.Get("user")
	if ok {
		username = user.(*model.User).Name
	}
	ctxIf, ok := c.Get("request-context")
	if !ok {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError,
			"error": "内部服务错误，请联系开发人员处理"})
		c.Abort()
		return
	}
	ctx := ctxIf.(context.Context)
	cancelIf, ok := c.Get("request-cancel")
	if !ok {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError,
			"error": "内部服务错误，请联系开发人员处理"})
		c.Abort()
		return
	}
	cancel := cancelIf.(context.CancelFunc)
	if cancel != nil {
		defer cancel()
	}
	timeout := false
	// 接收到请求时进行记录
	log.Printf("接收访问请求：%s，请求方法：%s，请求标识:%s，用户：%s",
		c.Request.URL.Path, c.Request.Method, requestID, username)
	defer func() {
		if timeout {
			log.Print("请求处理超时，请稍后再试")
			c.JSONP(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "error": "请求处理超时，请稍后再试"})
		} else {
			responseIf, ok := c.Get("response")
			if ok {
				response := responseIf.(*model.Response)
				if response.Err != nil {
					log.Print(response.Err)
					c.JSONP(http.StatusOK, gin.H{"code": response.Code, "error": errors.GetInfo(response.Err)})
				} else {
					c.JSONP(http.StatusOK, gin.H{"code": http.StatusOK, "total": response.Total, "spec": response.Spec})
				}
			} else {
				c.Status(http.StatusOK)
			}
		}
		// 请求处理完毕时也进行记录
		log.Printf("响应访问请求：%s，请求方法：%s，请求标识:%s，用户：%s",
			c.Request.URL.Path, c.Request.Method, requestID, username)
	}()
	done := make(chan struct{})
	go func() {
		c.Next()
		close(done)
	}()
	select {
	case <-ctx.Done():
		timeout = true
	case <-done:
		return
	}
}
