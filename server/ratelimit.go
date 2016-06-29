package server

import (
    "fmt"
	"github.com/gin-gonic/gin"
	"github.com/manucorporat/stats"
)

var ips = stats.New()

func rateLimit(c *gin.Context) {
    
	ip := c.ClientIP()
	value := int(ips.Add(ip, 1))

	if value >= 10 {
		if value % 10 == 0 {
			fmt.Println("blocked", ip)
		}
		c.Abort()
		c.String(503, "you were automatically banned :)")
	}
}