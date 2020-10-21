package taillog

import (
	"context"
	"fmt"
	"github.com/nxadm/tail"
	"logagent/kafka"
)


//TailTask: 一个日志收集的任务
type TailTask struct {
	Path string
	Topic string
	Instance *tail.Tail
	//为了实现推出t.run
	ctx context.Context
	cancelFunc context.CancelFunc
}

func NewTailTask(path , topic string)(tailObj *TailTask)  {
	ctx , cancel := context.WithCancel(context.Background())
	tailObj = &TailTask{
		Path:       path,
		Topic:      topic,
		ctx:        ctx,
		cancelFunc: cancel,
	}
	tailObj.init()	//根据路径去打开对应的日志
	return
}

func (t *TailTask)init()()  {
	config := tail.Config{
		ReOpen:	true,
		Follow:	true,
		Location: &tail.SeekInfo{Offset: 0,Whence: 2},
		MustExist: false,
		Poll:	true,
	}
	var err  error
	t.Instance,err = tail.TailFile(t.Path,config)
	if err != nil{
		fmt.Println("tail file failed,err%v\n",err)
	}
	go t.run() //直接去采集日志发送到kafka
}

func (t *TailTask)run()  {
	for {
		select {
		case <- t.ctx.Done():
			fmt.Printf("tail task：%s_%s 结束...\n",t.Path,t.Topic)
			return
		case line:= <- t.Instance.Lines://从tailObj的通道中一行一行的读取日志数据
			//3.2发往kafka
			//kafka.SendTokafka(t.Topic,line.Text)//函数调用函数
			//先把日志数据发到一个通道中
			kafka.SendToChan(t.Topic,line.Text)
			//kafka那个包中有单独的Goroutine去取日志数据发往kafka
		}
	}
}
