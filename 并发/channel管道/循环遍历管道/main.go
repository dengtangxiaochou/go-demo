package main

import "fmt"

func main()  {
	var ch  = make(chan int,10)
	for i := 1; i<=10 ; i++ {
		ch <- i
	}
	close(ch)	//关闭管道
	//循环遍历
	for val  := range ch{
		fmt.Println(val)
	}

	//for遍历管道的时候不需要关闭管道
	var ch2  = make(chan int,10)
	for j := 1; j<=10 ; j++ {
		ch2 <- j
	}
	for i :=0 ;i <10 ; i ++ {
		fmt.Println(<-ch2)
	}

}
