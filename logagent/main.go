package	main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"logagent/config"
	"logagent/etcd"
	"logagent/kafka"
	"logagent/taillog"
	"logagent/utils"
	"sync"
	"time"
)

//入口程序

var  (
	cfg = new(config.AppConf)
)

//func run()  {
////	1、读取日志
//	for {
//		select {
//			case line := <- taillog.ReadChan():
//				//	2.发送到Kafka
//				kafka.SendTokafka(cfg.KafkaConf.Topic,line.Text)
//		default:
//			time.Sleep(time.Second)
//		}
//	}
//}


func main()  {
	//0、加载配置文件
	//cfg, err := ini.Load("./config/config.ini")
	//fmt.Println(cfg.Section("kafka").Key("address"))
	//fmt.Println(cfg.Section("kafka").Key("topic"))
	//fmt.Println(cfg.Section("taillog").Key("path"))

	err := ini.MapTo(cfg,"./config/config.ini")
	if err != nil{
		fmt.Println("init  failed, err:%V\n",err)
	}

	//	1.初始化Kafka连接
	err = kafka.Init([]string{cfg.KafkaConf.Address},cfg.KafkaConf.ChanMaxSize)
	if err != nil{
		fmt.Println("init kafka failed, err:%V\n",err)
		return
	}
	fmt.Println("init kafka success")

	//2、初始化ETCD
	err = etcd.Init(cfg.EtcdConf.Address,time.Duration(cfg.EtcdConf.Timeout)*time.Second)
	if err != nil{
		fmt.Println("init etcd failed, err:%V\n",err)
		return
	}
	fmt.Println("init etcd success")

	//为了实现每一个logagent拉去自己独有的配置，所以要以自己的ip地址作为区分
	ipStr,err := utils.GetOutboundIP()
	if err != nil {
		panic(err)
	}
	etcdConfKey := fmt.Sprintf(cfg.EtcdConf.Key,ipStr)
	//2.1从etcd获取日志收集项的信息
	logEntryConf, err := etcd.GetConf(etcdConfKey)

	if err != nil {
		fmt.Printf(" etcd.GetConf ,err%v\n",err)
		return
	}
	fmt.Printf("get conf from etcd succes,%v\n",logEntryConf)
	//
	//2.2派一个哨兵去监视日志收集项的变化
	for index,value := range logEntryConf{
		fmt.Printf("index:%v value%v\n",index,value)
	}
	//3、收集日志发往Kafka
	taillog.Init(logEntryConf)	//因为Neconfchar访问了Tskm的newConfchan，这个channel是在taillog.Init(logEntryConf)执行的初始化
	NewConfChan := taillog.NewConfChan() //从taillog包中获取对外暴露的通道
	var wg sync.WaitGroup
	wg.Add(1)
	go etcd.WatchConf(etcdConfKey,NewConfChan) //哨兵发项最新的配置信息回通知上面的那个通道
	wg.Wait()

}


