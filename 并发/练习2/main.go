package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func putNum(inChan chan int)  {
	for i := 2; i< 1200000 ;i++ {
		inChan <- i
	}
	close(inChan)
	wg.Done()
}

func primeNum(inChan chan int,primChan chan int,exitChan chan bool)  {
	for num := range inChan{
		var flag  = true
		for i := 2; i< num ;i++  {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag{
			primChan <- num
		}
	}
	exitChan <- true
	wg.Done()
	//

}

//打印
func printPrime(primeChan chan int)  {
	for v := range primeChan{
		fmt.Println(v)
	}
	wg.Done()
}


func main()  {
	intChan := make(chan int,1000)
	primeChan := make(chan int,1000)
	exitChan := make(chan bool,16)
	//存放数据的协程
	wg.Add(1)
	go putNum(intChan)

	//统计素数的协程
	for i := 0 ;i < 16 ;i++ {
		wg.Add(1)
		go primeNum(intChan,primeChan,exitChan)
	}
	//打印素数协程
	wg.Add(1)
	go printPrime(primeChan)

	wg.Add(1)
	go func() {
		for i := 0 ; i < 16;i++  {
			<- exitChan
		}

		//关闭
		close(primeChan)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("完成")
}
