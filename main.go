package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-erp-api/bootstrap"
)

func main() {
	// new 一个 Gin Engine 实例
	router := gin.New()

	// 初始化路由绑定
	bootstrap.SetupRoute(router)

	// 运行服务
	err := router.Run(":3000")
	if err != nil {
		// 错误处理, 端口占用了或其他错误信息
		fmt.Println(err.Error())
	}
}
