// Package bootstrap 处理程序初始化逻辑
package bootstrap

import (
	"github.com/gin-gonic/gin"
	"go-erp-api/app/http/middlewares"
	"go-erp-api/routes"
	"net/http"
	"strings"
)

// SetupRoute 路由初始化
func SetupRoute(router *gin.Engine) {
	// 注册全局中间件
	registerGlobalMiddleWare(router)
	// 注册API路由
	routes.RegisterAPIRoutes(router)
	// 配置404路由
	setup404Handler(router)
}

// 注册全局中间件
func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		middlewares.Logger(),
		middlewares.Recovery(),
	)
}

// 配置404路由
func setup404Handler(router *gin.Engine) {
	// 处理 404 请求
	router.NoRoute(func(c *gin.Context) {
		// 获取标头信息的 Accept 信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是 HTML 的话
			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			// 默认返回 JSON
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})
}
