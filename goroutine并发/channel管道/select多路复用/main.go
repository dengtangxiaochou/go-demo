package main

import (
	"fmt"
	"time"
)

var ininChan = make(chan interface{},10)
var stringChan  =make(chan string,5)

func main()  {

	//再一个int类型管道内写入数据
	for i := 0; i< 10 ; i++ {
		ininChan <- i
	}

	//再一个string类型的管道写入数据
	for i := 0; i< 5 ; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d",i)
	}
	//使用多路复用select不需要关闭chan管道
	for {
		select {
		case v:= <- ininChan:
			fmt.Printf("从intChan 读取数据%v\n",v)
			time.Sleep(time.Millisecond * 30)
		case v := <- stringChan:
			fmt.Printf("从stringchan 读取数据%v\n",v)
			time.Sleep(time.Millisecond * 30)
		default:
			fmt.Printf("获取数据完毕")
			return
		}
	}

}
