package boot

import (
	"cn.sockstack/gin_demo/models"
	"cn.sockstack/gin_demo/pkg/config"
	"cn.sockstack/gin_demo/routers"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)
var Gin  = gin.Default()

func Run()  {
	panicHandle()
	//启动初始化
	bootInit()
	//设置运行模式
	setMode()

	//注册路由
	routers.Init(Gin)

	// 监听并在 0.0.0.0:8080 上启动服务
	Gin.Run(
		fmt.Sprintf(
			"%s:%d",
			config.Server.Address,
			config.Server.Port,
		),
	)
	close()
}

func bootInit()  {
	models.Init()
}

func setMode()  {
	if config.App.Mode == "debug" {
		gin.SetMode(gin.ReleaseMode)
		models.DB.LogMode(true)
	}
}

func close()  {
	models.Close()
}

func panicHandle()  {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()
}
