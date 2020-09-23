package main

import (
	"fmt"
	"time"
)

//函数
func sayHello()  {
	for i := 0;i <10 ; i++  {
		time.Sleep(time.Microsecond * 50)
		fmt.Println("Hello,world")
	}

}

//函数
func test()  {
	//这里我们用defer + recaover
	defer func() {
		//捕获test抛出的panic
		if err := recover();err != nil {
			fmt.Println("test()  发生错误",err)
		}
	}()
	var myMap map[int]string
	myMap[0] ="go"


}

func main()  {

	go sayHello()

	go test()

	time.Sleep(time.Second)


}