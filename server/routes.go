package server

import (
	"github.com/gin-gonic/gin"
)


func StartServer(){

    gin.SetMode(gin.ReleaseMode)

	s := gin.New()
	//set middleware rate limit
    s.Use(rateLimit, gin.Recovery())

	s.GET("/", index)
	//start server 
	s.Run(":8090")
}


func index(c *gin.Context) {
	c.String(200, "api gateway is running")
}


