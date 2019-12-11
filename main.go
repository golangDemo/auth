package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/mqtt/auth", auth)
	r.GET("/mqtt/acl", acl)
	r.Run(":9090")
}

func auth(c *gin.Context) {
	username := c.PostForm("username")
	clientid := c.PostForm("clientid")
	password := c.PostForm("password")

	fmt.Println("recv request ", username, clientid, password)
}

func acl(c *gin.Context) {
	username := c.PostForm("username")
	topic := c.PostForm("topic")
	access := c.PostForm("access")

	fmt.Println("recv request ", username, topic, access)
}
