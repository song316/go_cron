package main

import (
	"fmt"
	"github.com/song316/go_cron/cron"
)

func main() {
	//r := gin.Default()
	//r.POST("/", func(c *gin.Context) {
	//	name := c.DefaultPostForm("name", "ryan")
	//	age := c.DefaultPostForm("age", "1")
	//	c.JSON(200, gin.H{
	//		"name":name,
	//		"age":age,
	//	})
	//})
	//r.Run(":8080")

	flag := make(chan bool)
	c := cron.New();
	c.AddFunc("*/1 * * * * *", func() {
		fmt.Println("每秒")
	})
	c.Start()
	<-flag
}