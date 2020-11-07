package session

import (
	"errors"
	"github.com/garyburd/redigo/redis"
	uuid "github.com/satori/go.uuid"
	"sync"
	"time"
)

type RedisSessionMgr struct {
	//redis地址
	addr string
	//密码
	passwd string
	//连接池
	pool *redis.Pool
	//锁
	rwlock sync.RWMutex
	//大Map
	sessionMap map[string]Session
}

func (r *RedisSessionMgr) Init(addr string, options ...string) (err error) {
	//若有其他参数
	if len(options) > 0 {
		r.passwd = options[0]
	}
	//创建连接池
	r.pool =myPool(addr,r.passwd)
	r.addr = addr
	return
}

func myPool(addr, password string) *redis.Pool {
	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}
			//若有密码判断
			if _, err := conn.Do("AUTH", password); err != nil {
				conn.Close()
				return nil, err
			}
			return conn, err
		},
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			_, err := conn.Do("PING")
			return err
		},
		MaxIdle:         64,
		MaxActive:       1000,
		IdleTimeout:     240 * time.Second,
		Wait:            false,
		MaxConnLifetime: 0,
	}

}

func (r *RedisSessionMgr) CreateSession() (session Session, err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	//用UUID作为sessionID
	id := uuid.NewV4()
	//转string
	sessionId := id.String()
	//创建个Session
	session = NewRedisSession(sessionId,r.pool)
	//加入到大MAP
	r.sessionMap[sessionId] = session
	return
}

func (r *RedisSessionMgr) Get(sessionId string) (session Session, err error) {
	//加锁
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	//先判断内存
	session, ok := r.sessionMap[sessionId]
	if !ok {
		err = errors.New("session not exists")
		return
	}
	return
}

//构造函数
func NewRedisSessionMgr() SessionMgr {
	sr := &RedisSessionMgr{
		addr:       "",
		passwd:     "",
		pool:       nil,
		rwlock:     sync.RWMutex{},
		sessionMap: make(map[string]Session, 32),
	}
	return sr
}
