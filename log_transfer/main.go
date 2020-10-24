package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log_transfer/config"
	"log_transfer/es"
	"log_transfer/kafka"
)

var (
	cfg = new(config.LogTransfer)
)

//log transfer
//将日志数据从kafka取出来发往es
func main() {
	//0.加载配置文件
	err := ini.MapTo(cfg, "./config/cfg.ini")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("cfg:%v\n", cfg)
	//2.初始化ES
	err = es.Init(cfg.EsCfg.Address, cfg.EsCfg.ChanSize,cfg.EsCfg.Nums)
	if err != nil {
		fmt.Printf("init ES Clinet failed ,err:%v\n", err)
		return
	}
	fmt.Println("init es success.")
	//1.初始化Kafka连接
	//1.1连接Kafka，创建分区的消费者
	//1.2每个分区的消费者分别取出数据，通过SendToes（）将数据发往ES
	err = kafka.Init([]string{cfg.KafkaCfg.Address}, cfg.KafkaCfg.Topic)
	if err != nil {
		fmt.Printf("init Kafka consumer failed ,err:%v\n", err)
		return
	}
	select {}
	//2.从Kafka取数据
	//3.发往es
}
