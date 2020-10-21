package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

var (
	cli *clientv3.Client
)
//需要收集的日志的配置信息
type LogEn struct {
	Path string `json:"path"`	//日志存放的路径
	Topic string `json:"topic"`	//日志要fa发往kafka的tpoic
}

func Init(addr string,timeout time.Duration)(err error)  {
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{addr},
		DialTimeout: timeout,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	return
}

//从ETCD中根基KEy获取配置项目
func GetConf(key string)(logEntryConf []*LogEn,err error) {
	// get
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	for _, ev := range resp.Kvs {
		//fmt.Printf("%s:%s\n", ev.Key, ev.Value)
		err = json.Unmarshal(ev.Value, &logEntryConf)
		if err != nil {
			fmt.Println("unmarshal etcd value failed, err%v\n", err)
			return
		}
	}
	return
}

//etcd watch
func WatchConf(key string,newConfCh chan <- []*LogEn)  {
	// watch key:q1mi change
	rch := cli.Watch(context.Background(), key) // <-chan WatchResponse
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("Type: %s Key:%s Value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			//通知tallog。taskMgr
			//1.先判断操作类型
			var newConf []*LogEn
			if ev.Type != clientv3.EventTypeDelete{
				//如果是删除操作，手动传递一个空的配置项目
				err := json.Unmarshal(ev.Kv.Value,&newConf)
				if err != nil {
					fmt.Printf("unmarshal failed, err%v\n",err)
					continue
				}
			}
			fmt.Printf("Get New conf%v\n",newConf)
			newConfCh <- newConf
		}
	}
}