package main

import (
	"fmt"
)

//单向管道
func main()  {
	//1、默认是双向的
	ch1 := make(chan int,2)
	ch1 <- 10
	ch1 <- 12
	m1 := <-ch1
	m2 := <-ch1
	fmt.Println(m1,m2)

	//管道只写
	ch2 := make(chan<- int, 2)
	ch2 <- 10
	ch2 <- 19
	//<-ch2 灭法读取

	//管道只读
	ch3 := make(<-chan int,2 )
	//ch3 <- 10
	fmt.Println(ch3)

}