package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// InitRouter 初始化 HTTP 路由
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	//Propagators用b3即可（不区分大小写）， traceProvider使用全局自定义的
	apiGroup := r.Group("/api/")
	apiGroup.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"code": 0, "msg": "success", "data": "pong", "ctx": fmt.Sprintf("%#v", ctx)})
	})

	return r
}

func noTraceFilter(r *http.Request) bool {
	if r.URL.Path == "/metrics" || r.URL.Path == "/ping" {
		return false
	}
	return true
}
