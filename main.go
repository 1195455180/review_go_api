package main

import (
	"flag"
	"fmt"
	"gohub/bootstrap"
	btsConfig "gohub/config"
	"gohub/pkg/config"

	"github.com/gin-gonic/gin"
)

func init() {
	// 加载 config 目录下的配置信息
	btsConfig.Initialize()
}

func main() {

	// 配置初始化
	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()
	config.InitConfig(env)

	// 初始化 Gin 实例
	router := gin.New()

	// 初始化 DB
	bootstrap.SetupDB()

	// 初始化路由绑定
	bootstrap.SetRoute(router)

	err := router.Run(":" + config.Get("app.port"))

	if err != nil {
		// 错误处理
		fmt.Println(err.Error())
	}
}
