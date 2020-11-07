package main

import (
	"github.com/gin-gonic/gin"
	"go-demo/controller"
	"go-demo/dao/db"
)

func main() {
	roarer := gin.Default()
	dns := "root:root@tcp(192.168.2.32:3306)/blogger?parseTime=true"
	err := db.Init(dns)
	if err != nil {
		panic(err)
	}

	//加载静态文件
	roarer.Static("/static","./static")
	//加载模板
	roarer.LoadHTMLGlob("views/*")
	roarer.GET("/",controller.IndexHandler)
	roarer.GET("/category/",controller.CategoryList)
	_ = roarer.Run(":8080")
}
