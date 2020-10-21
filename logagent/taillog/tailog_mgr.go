package taillog

import (
	"fmt"
	"logagent/etcd"
	"time"
)


var tskMgr *tailLogMgr

//tailTASK管理者
type tailLogMgr struct {
	logEnty []*etcd.LogEn
	tskMap map[string]*TailTask
	newConfChan chan []*etcd.LogEn
}



func Init(logEntryConf []*etcd.LogEn)  {
	tskMgr = &tailLogMgr{
		logEnty: logEntryConf,	//把当前的日志收集项存储起来
		tskMap: make(map[string]*TailTask,16),
		newConfChan: make(chan []*etcd.LogEn),
	}
	for _, logEnty := range logEntryConf{
		//conf：etcd.LoFenty
		//logEnty.Path:要收集的日志路径
		//初始化的时候起了多少个tailtask 都要记下来，为了后续判断方便
		tailObj := NewTailTask(logEnty.Path,logEnty.Topic)
		mk := fmt.Sprintf("%s_%s",logEnty.Path,logEnty.Topic)
		tskMgr.tskMap[mk] = tailObj
	}
	go tskMgr.run()
}

//监听自己的NewConfchan 有了新的配置过来以后做对应的处理

func (t *tailLogMgr)run()  {
	for {
		select {
		case newConf := <- t.newConfChan:
			for _, conf := range newConf{
				mk := fmt.Sprintf("%s_%s",conf.Path,conf.Topic)
				_, ok := t.tskMap[mk]
				if ok {
					//原来就有，不需要操作
					continue
				}else {
					//新增的
					tailObj := NewTailTask(conf.Path,conf.Topic)
					t.tskMap[mk] = tailObj
				}
			}
			//找出原来t.logEnty有，但是newconf没有的，要删掉
			for  _, c1 := range t.logEnty{
				isDelete := true
				for _, c2 := range newConf{
					if c2.Path == c1.Path && c2.Topic == c1.Topic{
						isDelete = false
						continue
					}
				}
				if isDelete{
					//把c1对应的这个TailObj给停掉
					mk := fmt.Sprintf("%s_%s",c1.Path,c1.Topic)
					t.tskMap[mk].cancelFunc()
				}
			}
			//2.配置删除
			fmt.Println("新的配置来了！",newConf)
			default:
			time.Sleep(time.Second)
		}
	}
}

//向外暴露一个函数，向tskMgr的newConfChan
func NewConfChan() chan <- []*etcd.LogEn{
	return tskMgr.newConfChan
}