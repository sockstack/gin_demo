package routers

import "github.com/gin-gonic/gin"

func Init(r *gin.Engine)  {
	//注册test路由
	test(r)
}