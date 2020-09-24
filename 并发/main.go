package main

import (
	"fmt"
	"sync"
	"time"
)

var wg  sync.WaitGroup

func tert2()  {
	start := time.Now().Unix()
	for num :=2 ; num <120000 ; num++  {
		var flag1  = true
		for i :=2 ;i<num ; i++  {
			if num%i == 0 {
				flag1 = false
				break
			}
		}
		if flag1 {
			fmt.Println(num,"是素数")
		}
	}
	end := time.Now().Unix()
	fmt.Println(end-start)
	wg.Done()
}



func test()  {
	for i := 0 ;i < 10 ; i++  {
		fmt.Println("你好test",i)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done()
}


func main()  {
	wg.Add(1)
	 go test() //开启一个协程
	for i := 0 ;i < 10 ; i++  {
		fmt.Println("你好",i)
		time.Sleep(time.Millisecond * 50)
	}
	wg.Wait()
	fmt.Println("主进程退出。。。。")

	for i := 1; i<=5 ;i++  {
		wg.Add(1)
		go tert2()
	}
	wg.Wait()
	fmt.Println("主进程退出。。。。")
}
