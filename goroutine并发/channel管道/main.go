package main

import "fmt"

func main() {
	//创建channel
	ch := make(chan int, 3)
	//给管道赋值
	ch <- 10
	ch <- 21
	ch <- 30
	//获取channel的值
	a := <-ch
	fmt.Println(a) //10

	<-ch //21
	c := <-ch
	fmt.Println(c) //3 0

	ch <- 52
	ch <- 60
	fmt.Printf("值：%v,容量%v,长度%v", ch, cap(ch), len(ch))

	//管道的类型
	ch1 := make(chan int, 4)
	ch1 <- 14
	ch1 <- 15
	ch1 <- 16

	ch2 := ch1
	ch2 <- 25
	<-ch1
	<-ch1
	<-ch1
	d := <-ch1
	fmt.Println(d)

	//ch6 := make(chan int, 1)
	//ch6 <- 34
	//ch6 <- 55

}