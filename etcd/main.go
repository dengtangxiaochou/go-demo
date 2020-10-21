package main

import (
	"context"
	"fmt"
	"time"
	"go.etcd.io/etcd/clientv3"
)

// etcd client put/get demo
// use etcd/clientv3

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()
	// put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	value := `[{"path":"./nginx.log","topic":"web_log"},{"path":"./my.log","topic":"redis_log"}]`
	_, err = cli.Put(ctx, "/logagent/10.154.69.197/collect_config", value)
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}
	//// get
	//ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	//resp, err := cli.Get(ctx, "q1mi")
	//cancel()
	//if err != nil {
	//	fmt.Printf("get from etcd failed, err:%v\n", err)
	//	return
	//}
	//for _, ev := range resp.Kvs {
	//	fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	//}
}

//
//package main
//
//import (
//	"context"
//	"fmt"
//	"time"
//
//	"go.etcd.io/etcd/clientv3"
//)
//
//// watch demo
//
//func main() {
//	cli, err := clientv3.New(clientv3.Config{
//		Endpoints:   []string{"127.0.0.1:2379"},
//		DialTimeout: 5 * time.Second,
//	})
//	if err != nil {
//		fmt.Printf("connect to etcd failed, err:%v\n", err)
//		return
//	}
//	fmt.Println("connect to etcd success")
//	// get
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
//	resp, err := cli.Get(ctx, "q1mi")
//	cancel()
//	if err != nil {
//		fmt.Printf("get from etcd failed, err:%v\n", err)
//		return
//	}
//	for _, ev := range resp.Kvs {
//		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
//	}
//
//	defer cli.Close()
//	// watch key:q1mi change
//	rch := cli.Watch(context.Background(), "q1mi") // <-chan WatchResponse
//	for wresp := range rch {
//		for _, ev := range wresp.Events {
//			fmt.Printf("Type: %s Key:%s Value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
//		}
//	}
//
//}