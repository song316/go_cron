package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		name := c.DefaultPostForm("name", "ryan")
		age := c.DefaultPostForm("age", "1")
		c.JSON(200, gin.H{
			"name":name,
			"age":age,
		})
	})
	r.Run(":8080")
}