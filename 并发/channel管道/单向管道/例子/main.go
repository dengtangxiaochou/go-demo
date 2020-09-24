package main

import (
	"fmt"
	"sync"
	"time"
)
var wg = sync.WaitGroup{}
var ch = make(chan int,10)

func fn1(ch chan <- int)  {
	for n := 1;n <= 10 ;n++  {
		ch <- n
		fmt.Printf("【写入】数据%v成功\n",n)
		time.Sleep(time.Millisecond * 30)
	}
	close(ch)
	wg.Done()
}

func fn2(ch <- chan int)  {
	for v := range ch{
		//fmt.Println(v)
		fmt.Printf("【读取】数据%v成功\n",v)
		time.Sleep(time.Millisecond * 30)
	}

	wg.Done()
}

func main()  {
	wg.Add(1)
	go fn1(ch)

	wg.Add(1)
	go fn2(ch)
	wg.Wait()
	fmt.Println("成功")
}


