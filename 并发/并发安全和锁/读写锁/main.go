package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

var mutest sync.RWMutex
//写方法
func write()  {
	mutest.Lock()
	fmt.Println("执行写操作")
	time.Sleep(time.Second * 2)
	mutest.Unlock()
	wg.Done()
}

//读方法
func read()  {
	mutest.RLock()
	fmt.Println("--执行读操作")
	time.Sleep(time.Second * 2)
	mutest.RUnlock()
	wg.Done()
}

func main()  {
	//开启10个协程执行写
	for i := 0; i <10 ; i++ {
		wg.Add(1)
		go read()
	}

	//开启10个协程执行读
	for i := 0; i <10 ; i++ {
		wg.Add(1)
		go write()
	}
	wg.Wait()
}

