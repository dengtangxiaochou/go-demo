package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取客户端cookIE并且校验
		cookie, err := c.Cookie("abc")
		if err != nil {
			if cookie == "123" {
				c.Next()
				return
			}
		}
		//返回错误信息
		c.JSON(http.StatusUnauthorized, gin.H{"error": "err"})
		//如果验证不通过，不在调用后续的函数
		c.Abort()
		return
	}
}

func main() {
	r := gin.Default()

	r.GET("/login", func(c *gin.Context) {
		//设置COOKIE
		c.SetCookie("abc", "123", 60, "/",
			"localhost", false, true)
		c.String(200, "Login success")
	})

	r.GET("/home", AuthMiddleWare(), func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "home"})
	})
	_ = r.Run(":8080")
}
