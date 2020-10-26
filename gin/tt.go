package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//定义接收数据的结构体
type Login struct {
	//binding:"required"修饰的字段。若接收为空值，则报错，是必选字段
	User string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func main()  {
	r := gin.Default()
	//json绑定
	r.POST("loginJSON", func(c *gin.Context) {
		//声明接收的变量
		var json Login
		//将requers的boby中的数据，自动按照Json格式解析到结构体
		err := c.ShouldBindJSON(&json)
		if err != nil {
			//gin.H封装了生成了Json数据的工具
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}
		//判断用户名密码是否正确
		if json.User != "root" || json.Password != "admin" {
			c.JSON(http.StatusBadRequest,gin.H{"status":"304"})
			return
		}
		c.JSON(http.StatusOK,gin.H{"status":"200"})
	})

	//表单
	r.POST("loginForm", func(b *gin.Context) {
		var form Login
		//Bind()默认解析并绑定Form格式
		//根据请求头Content-type自动推断
		err := b.Bind(&form)
		if err != nil {
			b.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}
		//判断用户名密码是否正确
		if form.User != "root" || form.Password != "admin" {
			b.JSON(http.StatusBadRequest,gin.H{"status":"304"})
			return
		}
		b.JSON(http.StatusOK,gin.H{"status":"200"})
	})

	//url
	r.GET("/:user/:password", func(b *gin.Context) {
		var url Login
		//Bind()默认解析并绑定Form格式
		//根据请求头Content-type自动推断
		err := b.ShouldBindUri(&url)
		if err != nil {
			b.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}
		//判断用户名密码是否正确
		if url.User != "root" || url.Password != "admin" {
			b.JSON(http.StatusBadRequest,gin.H{"status":"304"})
			return
		}
		b.JSON(http.StatusOK,gin.H{"status":"200"})
	})
	_ = r.Run(":8000")
}