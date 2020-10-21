package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

type logData struct {
	topic string
	data string
}

//
var (
	client sarama.SyncProducer
	logDataChan chan *logData
)

func Init(addrs []string,maxSize int)(err error)  {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 连接kafka
	client, err = sarama.NewSyncProducer(addrs, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	fmt.Println("连接Kafka成功")
	logDataChan = make(chan *logData,maxSize)
	//开启后台的goroutine从通道中读取数据发往Kafka
	go sendTokafka()
	return
}
//给外部暴露的一个函数，该函数只把日志数据发送到一个内部的Channel中
func SendToChan(topic,data string)  {
	msg := &logData{
		topic: topic,
		data: data,
	}
	logDataChan <- msg
}

//正真往kafka发送日志的函数
func sendTokafka() {
	for {
		select {
		case ld :=  <- logDataChan:
			// 构造一个消息
			msg := &sarama.ProducerMessage{}
			msg.Topic = ld.topic
			msg.Value = sarama.StringEncoder(ld.data)
			// 发送消息
			pid, offset, err := client.SendMessage(msg)
			if err != nil {
				fmt.Println("send msg failed, err:", err)
				return
			}
			fmt.Printf("pid:%v offset:%v\n", pid, offset)
		default:
			time.Sleep(time.Millisecond*50)
		}
	}
}