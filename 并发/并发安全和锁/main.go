package main

import (
	"fmt"
	"sync"
	"time"
)
//
//var count  = 0

var wg  sync.WaitGroup

var mutest sync.Mutex

var m1 = make(map[int]int,0)

//func test()  {
//	mutest.Lock()
//	count++
//	fmt.Println("the count is :",count)
//	time.Sleep(time.Millisecond)
//	mutest.Unlock()
//	wg.Done()
//}

func test2(num int) {
	mutest.Lock()
	var sum = 1
	for i := 1;i <= num ; i++ {
		sum *= i
	}
	m1[num] = sum
	fmt.Printf("key=%v value=%v\n",num,sum)
	time.Sleep(time.Millisecond)
	mutest.Unlock()
	wg.Done()
}
func main()  {
	for  r := 0;r < 60  ; r++  {
		//wg.Add(1)
		//go test()
		wg.Add(1)
		go test2(r)

	}
	//for  n := 0;n < 40  ; n++  {
	//	wg.Add(1)
	//	go test2(n)
	//
	//}
	wg.Wait()
}
