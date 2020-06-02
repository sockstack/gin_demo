package main
// 导入gin包
import (
	"cn.sockstack/gin_demo/pkg/config"
	"cn.sockstack/gin_demo/routers"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 入口函数
func main() {
	// 初始化一个http服务对象
	r := gin.Default()

	//注册路由
	routers.Init(r)

	r.Run(fmt.Sprintf("%s:%d", config.Server.Address, config.Server.Port)) // 监听并在 0.0.0.0:8080 上启动服务
}