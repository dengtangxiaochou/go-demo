package session

import "fmt"

//中间件让用户选择使用的版本
var (
	sessionMgr SessionMgr
	)


func Init(provider string,addr string,options ...string)(err error)  {
	switch provider {
	case "memory":
		sessionMgr = NewMemorySeesionMgr()
	case "redis":
		sessionMgr = NewRedisSessionMgr()
	default:
		fmt.Println("不支持")
		return
	}
	err = sessionMgr.Init(addr, options...)
	return
}