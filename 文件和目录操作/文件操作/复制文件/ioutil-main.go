package main

import (
	"fmt"
	"io/ioutil"
)

func copy(srcFileName string,dstFileName string) error  {
	byte, err  := ioutil.ReadFile(srcFileName)
	if err != nil {
		return err
	}
	eer2  := ioutil.WriteFile(dstFileName,byte,0666)
	if eer2 != nil{
		return err
	}
	return nil
}
func main()  {
	src := "F:/test.txt"
	dst := "D:/test.txt"
	err := copy(src,dst)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("复制文件")
}
