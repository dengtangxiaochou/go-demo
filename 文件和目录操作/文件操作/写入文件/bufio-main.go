package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	file, err  := os.OpenFile("F:/test.txt",os.O_CREATE|os.O_RDWR,0666)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	writer := bufio.NewWriter(file)
	writer.WriteString("你好")
	writer.Flush()
}
