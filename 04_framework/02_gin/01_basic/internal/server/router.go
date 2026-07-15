package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine{
	r:=gin.Default()
	gin.SetMode(gin.ReleaseMode)

	if err := r.SetTrustedProxies(nil); err != nil {
		panic(err)
	}

	r.GET("/health",func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"ok":true,
			"status":"healthy",

		})
	})

	return r;
}