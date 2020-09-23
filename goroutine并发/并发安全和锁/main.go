package main

import (
	"fmt"
	"sync"
	"time"
)

var count  = 0

var wg  sync.WaitGroup

var mutest sync.Mutex

func test()  {
	mutest.Lock()
	count++
	fmt.Println("the count is :",count)
	time.Sleep(time.Millisecond)
	mutest.Unlock()
	wg.Done()
}

func main()  {
	for  r := 0;r < 20  ; r++  {
		wg.Add(1)
		go test()
	}
	wg.Wait()
}
