package main

import "fmt"

type Usber interface {
	start()
	stop()
}

//电脑
type Computer struct {
}

func (c Computer) work(usb Usber)  {
	//判断usb类型
	switch usb.(type) {
	case Phone:
		usb.start()
	case Camera:
		usb.stop()
	}
}

//手机

type Phone struct {
	Name string
}

func (p Phone) start()  {
	fmt.Println(p.Name,"启动")
}
func (p Phone)stop()  {
	fmt.Println(p.Name,"关闭")
}

//相机
type Camera struct {}

func (d Camera) start()  {
	fmt.Println("启动")
}
func (d Camera) stop()  {
	fmt.Println("关闭")
}


func main()  {
	var computer  = Computer{}
	var phone  = Phone{Name:"小米手机"}
	var camera  = Camera{}
	computer.work(phone)
	computer.work(camera)

	var p1  Usber=phone
	p1.start()

	var p3  = &Phone{Name:"华为"}
	var p4  Usber=p3
	p4.start()
}
