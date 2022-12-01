// Package auth 认证支持中间件
package auth

import (
	"net/http"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"transformer_mz/databases/connector"
	"transformer_mz/databases/model"
	"transformer_mz/libs/errors"
	"transformer_mz/libs/log"
	"transformer_mz/libs/token"
	"transformer_mz/middleware/trace"
)

// whiteList 无需认证token的地址
var whiteList = map[string]bool{
	"/v1/user/token": true,
}

func JWTAuth(c *gin.Context) {
	if reflect.ValueOf(c.Handler()).Pointer() == reflect.ValueOf(trace.Trace).Pointer() {
		c.Abort()
		return
	}
	// 设置开始执行的时间
	c.Set("begin", time.Now())
	inWhiteList := whiteList[c.Request.URL.Path]
	var err error
	// 从header获取token信息
	tokenStr := c.GetHeader("Token")
	if tokenStr == "" {
		// 从cookie获取token信息
		tokenStr, err = c.Cookie("Token")
		if err != nil {
			if !inWhiteList {
				c.AbortWithStatusJSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError,
					"error": "token获取失败，请重新登录"})
				c.Abort()
				return
			}
		}
		if tokenStr == "" && !inWhiteList {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"code": http.StatusUnauthorized,
				"error": "token无效，请重新登录"})
			c.Abort()
			return
		}
	}
	if tokenStr != "" {
		// 验证token有效性
		data, err := token.Validate(tokenStr)
		if err != nil {
			log.Print(err)
			if !inWhiteList {
				c.AbortWithStatusJSON(http.StatusOK, gin.H{"code": http.StatusUnauthorized,
					"error": errors.GetInfo(err)})
				c.Abort()
				return
			}
		}
		if err == nil {
			// 根据token中的用户编号查找数据库获取用户完整信息
			user := &model.User{}
			err = errors.New(connector.DataSource.Where("id=?", data.ID).First(user).Error)
			if err != nil {
				log.Error(err)
				if !inWhiteList {
					c.AbortWithStatusJSON(http.StatusOK, gin.H{"code": http.StatusUnauthorized,
						"error": "用户信息查找失败，请检查"})
					c.Abort()
					return
				}
			} else {
				c.Set("user", user)
			}
		}
	}
}
