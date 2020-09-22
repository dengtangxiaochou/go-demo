package main

import (
	"fmt"
	"time"
)
/*
2006 年
01 月
02 日
03 时 21小时制   15 24小时制
04 分
05 秒
*/
func main()  {
	timeObj := time.Now()

	var str = timeObj.Format("2006-01-02 03:04:05")
	var str1 = timeObj.Format("2006-01-02 15:04:05")
	fmt.Println(timeObj)
	//12小时
	fmt.Println(str)
	//24小时制
	fmt.Println(str1)
    //获取时间戳
	uninxtime := timeObj.Unix()
	fmt.Println( "当前时间戳：",uninxtime)
	//时间戳转换成日期符号
	timerOBJ := time.Unix(uninxtime,0)
	var str2 = timerOBJ.Format("2006-01-02 15:04:05")
	fmt.Println(str2)

	//日期转换成时间戳
	var str3  =  "2020-09-19 14:54:03"
	var tmp  ="2006-01-02 15:04:05"

	timeObj2, _ := time.ParseInLocation(tmp,str3,time.Local)
	fmt.Println(timeObj2)
	fmt.Println(timeObj2.Unix())
	
	/*
	定时器
	*/
	
	ticker := time.NewTicker(time.Second)
    n := 5
	for t := range ticker.C {
		n--
		fmt.Println(t)
		if n == 0 {
			ticker.Stop()
			break
		}
	}

}
