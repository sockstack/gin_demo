package routers

import (
	"cn.sockstack/gin_demo/api"
	"github.com/gin-gonic/gin"
)

func test(r *gin.Engine)  {
	//定义/test路由
	r.GET("/test", api.Test)
}
