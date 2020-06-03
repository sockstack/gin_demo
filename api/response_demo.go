package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var data = gin.H{"code": http.StatusOK, "data": "test"}

func Json(c *gin.Context)  {

	c.JSON(http.StatusOK, data)
}

func Xml(c *gin.Context)  {
	c.XML(http.StatusOK, data)
}
