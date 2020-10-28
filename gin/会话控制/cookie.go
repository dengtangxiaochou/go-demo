package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main()  {
	r := gin.Default()

	r.GET("cookic", func(c *gin.Context) {
		//获取客户端是否携带cookie
		cookie ,err := c.Cookie("key_cookie")
		if err != nil {
			cookie = "NotSet"
			//设置cookie
			//max age int 单位为秒
			//path cokie 所在的目录
			//domain string 域名
			//secure 是否通过https访问
			// httpOnly bool 是否允许别人通过JS获取自己的 COOKIE
			c.SetCookie("key_cookie","value_cookie",60,"/",
				"localhost",false,true)
		}
		fmt.Println(cookie)
	})
	r.Run(":8080")
}
