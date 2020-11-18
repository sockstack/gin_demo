package boot

import (
	"cn.sockstack/gin_demo/docs"
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
	//注册路由
	routers.Init(Gin)
	models.Init()
	swaggerInit()
}

func setMode()  {
	if config.App.Mode == "debug" {
		gin.SetMode(gin.ReleaseMode)
		models.DB.LogMode(true)
	}
}

func swaggerInit()  {
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%d", config.Server.Address, config.Server.Port)
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
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
