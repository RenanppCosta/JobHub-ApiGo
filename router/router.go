package router

import "github.com/gin-gonic/gin"

func InitializerRouter(){
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	
	r.Run()
}

