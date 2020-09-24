package main

import (
	"fmt"
	"os"
)

func main()  {
	file, err  := os.OpenFile("F:/test.txt",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}
	////写入文件
	//for i :=0 ; i < 10 ; i++  {
	//	file.WriteString("直接写入"+strconv.Itoa(i)+"\r\n")
	//}
	var str = "写入"
	file.Write([]byte(str))

}
