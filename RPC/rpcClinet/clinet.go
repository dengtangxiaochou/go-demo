package main

import (
	"fmt"
	"log"
	"net/rpc"
)

//传的参数
type Params struct {
	Width, Height int
}


//主函数

func main()  {
	//面积
	//1.连接远程的rpc服务
	client, err := rpc.DialHTTP("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	//2.调用方法
	ret := 0
	err2 := client.Call("Rect.Area", Params{50, 100}, &ret)
	if err2 != nil {
	 log.Fatal(err2)
	}
	fmt.Println("面积：",ret)
	//周长
	err3 := client.Call("Rect.Perimeter", Params{50, 100}, &ret)
	if err3 != nil {
		log.Fatal(err3)
	}
	fmt.Println("周长：",ret)
}