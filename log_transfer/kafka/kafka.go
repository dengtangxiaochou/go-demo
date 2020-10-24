package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log_transfer/es"
)

//LogData

//初始话Kafka
//Init 初始化Client
func Init(adders []string,topic string) error {
	consumer,err := sarama.NewConsumer(adders,nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	partitionList, err := consumer.Partitions(topic) // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return err
	}
	fmt.Println(partitionList)
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition("web_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return err
		}
		//defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
				//直接发给ES
				ld := es.LogDate{Topic:topic,Data:string(msg.Value)}
				if err != nil {
					fmt.Println(err)
					continue
				}
				//_ = es.SendToES(topic, ld)//函数调用函数
				es.SendToESChan(&ld)
			}
		}(pc)
	}
	return err
}
