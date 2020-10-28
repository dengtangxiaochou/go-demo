package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

//定义中间件
func MiddleWare()gin.HandlerFunc  {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行")
		//设置变量
		c.Set("request","中间件")
		//执行函数
		c.Next()
		//中间件执行的
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕",status)
		t2 := time.Since(t)
		fmt.Println("time：",t2)
	}
}



func main()  {
	r := gin.Default()
	//注册中间件
	r.Use(MiddleWare())
	//为了规范
	{

		r.GET("/middleware", func(c *gin.Context) {
		//取值
		req,_ := c.Get("request")
		fmt.Println("request",req)
		//页面接受
		c.JSON(200,gin.H{"request":req})
		})
	}
	//根路由后面 是定义的局部的中间件
	r.GET("/middleware2",MiddleWare(),func(c *gin.Context) {
		//取值
		req,_ := c.Get("request")
		fmt.Println("request",req)
		//页面接受
		c.JSON(200,gin.H{"request":req})
	})

	err := r.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
