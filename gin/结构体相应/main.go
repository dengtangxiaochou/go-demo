package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
)

func main()  {
	r := gin.Default()
	//json
	r.GET("/someJson", func(d *gin.Context) {
		d.JSON(200,gin.H{"message":"someJson","status":200})
	})
	//结构体响应
	r. GET("/someStruct", func(c *gin.Context) {
		var msg struct{
			Name string
			Message string
			Number int
		}
		msg.Name="root"
		msg.Message="message"
		msg.Number=123
		c.JSON(200,msg)
	})
	//XML
	r.GET("/someXml", func(c *gin.Context) {
		c.XML(200,gin.H{"message":"abc"})
	})
	//yaml
	r.GET("/someYmal", func(c *gin.Context) {
		c.YAML(200,gin.H{"name":"zhangs"})
	})
	//protobuf格式
	r.GET("/someProtobuf", func(c *gin.Context) {
		reps := []int64{int64(1),int64(2)}
		//定义数据
		label := "label"
		//传protobuf格式数据
		data :=  &protoexample.Test{
			Label:            &label,
			Type:             nil,
			Reps:             reps,
			Optionalgroup:    nil,
			XXX_unrecognized: nil,
		}
		c.ProtoBuf(200,data)
	})

	_ = r.Run(":8001")
}
