// Package cors 跨域支持中间件
package cors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CORS struct {
	engine *gin.Engine
}

func (cors *CORS) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 允许前端传递以下header给后端
	w.Header().Add("Access-Control-Allow-Headers", "X-Requested-With, Authorization, Content-Type, "+
		"Cache-Control, ETag, content-range, Timeout, Deadline, Use-Cache, Token")
	// 允许前端使用任意域名访问后端
	w.Header().Add("Access-Control-Allow-Origin", "*")
	// 允许前端获取response header中的X-Request-Id字段
	w.Header().Add("Access-Control-Expose-Headers", "X-Request-Id")
	// 允许前端使用以下方法访问后端
	w.Header().Add("Access-Control-Allow-Methods", "HEAD, OPTIONS, GET, PUT, POST, PATCH, DELETE")
	// 对于OPTIONS方法直接返回
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	cors.engine.ServeHTTP(w, r)
}

func New(engine *gin.Engine) *CORS {
	return &CORS{engine}
}
