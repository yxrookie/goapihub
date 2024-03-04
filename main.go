package main

import (
	"flag"
	"fmt"
	"goapihub/bootstrap"
	"goapihub/config"
	pkgconfig "goapihub/pkg/config"

	"github.com/gin-gonic/gin"
)

func init () {
	config.Initialize()
}

func main() {

       // 配置初始化，依赖命令行 --env 参数
	   var env string
	   flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	   flag.Parse()
	   pkgconfig.InitConfig(env)
   
	   // new 一个 Gin Engine 实例
	   router := gin.New()
   
	   // 初始化路由绑定
	   bootstrap.SetupRoute(router)
   
	   // 运行服务
	   err := router.Run(":" + pkgconfig.Get("app.port"))
	   if err != nil {
		   // 错误处理，端口被占用了或者其他错误
		   fmt.Println(err.Error())
	   }
}