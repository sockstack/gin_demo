package routers

import (
	"cn.sockstack/gin_demo/api"
	"github.com/gin-gonic/gin"
)

func test(r *gin.Engine)  {
	//定义/test路由
	r.GET("/test", api.Test)

	//响应json数据
	r.GET("/json", api.Json)
	//响应xml数据
	r.GET("/xml", api.Xml)
}
