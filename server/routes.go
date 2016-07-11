package server

import (
	"os"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
)

func StartServer() {

    godotenv.Load()

	gin.SetMode(gin.ReleaseMode)

	s := gin.New()
	//set middleware rate limit
	s.Use(rateLimit, gin.Recovery())

	s.GET("/", index)
	
	//start server
	s.Run(os.Getenv("PORT"))
}

func index(c *gin.Context) {
	c.String(200, "api gateway is running")
}
