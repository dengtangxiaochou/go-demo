package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main()  {
	r := gin.Default()

	r.GET("/long_async", func(c *gin.Context) {
		//需要搞一个副本
		copyContext := c.Copy()
		//异步处理
		go func() {
			time.Sleep(3 * time.Second)
			log.Println("异步执行："+copyContext.Request.URL.Path)
		}()
	})
	//同步
	r.GET("long_sync", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		log.Println("同不执行："+c.Request.URL.Path)
	})
	err := r.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
