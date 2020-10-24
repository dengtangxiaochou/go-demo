package es

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"strings"
	"time"
)

type LogDate struct {
	Topic string `json:"topic"`
	Data  string `json:"data"`
}

var (
	client *elastic.Client
	ch     chan *LogDate
)

func Init(adders string, chanSize int, nums int) (err error) {
	if !strings.HasPrefix(adders, "http://") {
		adders = "http://" + adders
	}
	//初始化连接得到一个Clinet
	client, err = elastic.NewClient(elastic.SetURL(adders))
	if err != nil {
		// Handle error
		return
	}
	fmt.Println("connect to es success")
	ch = make(chan *LogDate, chanSize)
	for i := 0; i < nums; i++ {
		go SendToES()
	}

	return
}

//发送数据到ES
func SendToESChan(msg *LogDate) {
	ch <- msg
}

func SendToES() {
	for {
		select {
		case msg := <-ch:
			put1, err := client.Index().Index(msg.Topic).BodyJson(msg).Do(context.Background())
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
		default:
			time.Sleep(time.Second)
		}
	}
}
