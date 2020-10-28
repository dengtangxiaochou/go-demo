package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// gin的helloWorld

func main() {
	// 1.创建路由
	r := gin.Default()
	//路由组1 ， 处理Get请求
	v1 := r.Group("v1")
	{
		v1.GET("/login", login)
		v1.GET("/submit", submit)
	}

	v2 := r.Group("v2")
	{
		v2.POST("/login", login)
		v2.POST("/submit", submit)
	}

	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	//r.GET("/", func(c *gin.Context) {
	//	c.String(http.StatusOK, "hello World!")
	//})
	// 3.监听端口，默认在8080
	//r.GET("/user/:name/*action1", func(c *gin.Context) {
	//	name := c.Param("name")
	//	action := c.Param("action1")
	//	c.String(http.StatusOK, name+" is"+action)
	//})
	//r.GET("welcome", func(c *gin.Context) {
	//	name := c.DefaultQuery("name","Jack")
	//	c.String(http.StatusOK,fmt.Sprintf("Hellok %s",name))
	//})

	//限制删上传大小 8mb 默认为32MB
	//r.MaxMultipartMemory = 8 << 20
	//r.POST("/upload", func(c *gin.Context) {
	//	////表单取文件
	//	//file, _ := c.FormFile("file")
	//	//log.Println(file.Filename)
	//	////传到项目根目录，名字就用本身
	//	//c.SaveUploadedFile(file,file.Filename)
	//	////打印信息
	//	//c.String(200,fmt.Sprintf("%s upload",file.Filename))
	//
	//	form,err := c.MultipartForm()
	//	if err != nil {
	//		c.String(http.StatusBadRequest,fmt.Sprintf("get err %s",err.Error()))
	//	}
	//	//获取所有图片
	//	files := form.File["files"]
	//	//遍历所有的图片
	//	for _,file := range files {
	//		//逐个存储
	//		if err := c.SaveUploadedFile(file, file.Filename); err != nil {
	//			c.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
	//			return
	//		}
	//	}
	//	c.String(200,fmt.Sprintf("upload ok %d files",len(files)))
	////	//表单参数设置默认值
	////	type1 := c.DefaultPostForm("type", "alert")
	////	//接收其他的
	////	username := c.PostForm("username")
	////	password := c.PostForm("password")
	////	//多选
	////	hobby := c.PostFormArray("hobby")
	////	c.String(http.StatusOK,
	////		fmt.Sprintf("type is %s, uaername is %s, password is %s,hobby is %v",
	////			type1, username, password, hobby))
	//})
	_ = r.Run(":8000")
}

func login(c *gin.Context) {
	name := c.DefaultQuery("name","jack")
	c.String(200,fmt.Sprintf("hallo %s\n",name))
}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name","lily")
	c.String(200,fmt.Sprintf("hallo %s\n",name))
}
