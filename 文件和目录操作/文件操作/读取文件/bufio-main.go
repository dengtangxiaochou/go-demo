package 读取文件

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main()  {
	file, err  := os.Open("F:/go-demo/time/main.go")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}
	//bufio
	var fileStr string
	reader := bufio.NewReader(file)

	for  {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			fileStr += str
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		//fmt.Println(str)
		fileStr += str
	}
	fmt.Println(fileStr)
}
