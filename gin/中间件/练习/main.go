package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func myTime(c *gin.Context) {
	start := time.Now()
	c.Next()
	//统计时间
	since := time.Since(start)
	fmt.Println("程序用时", since)
}

func main() {
	r := gin.Default()

	r.Use(myTime)
	shoppingGroup := r.Group("/shopping")
	{
		shoppingGroup.GET("/index", shopIndexHandler)
		shoppingGroup.GET("/home", shopHomeHandler)
	}

	err := r.Run(":8081")
	if err != nil {
		fmt.Println(err)
	}
}

func shopIndexHandler(c *gin.Context) {
	time.Sleep(5 * time.Second)
}

func shopHomeHandler(c *gin.Context) {
	time.Sleep(3 * time.Second)
}
