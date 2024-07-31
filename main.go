package main

import (
	"fmt"
	"gohub/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化 Gin 实例
	router := gin.New()

	bootstrap.SetRoute(router)

	err := router.Run(":3000")

	if err != nil {
		// 错误处理
		fmt.Println(err.Error())
	}
}
