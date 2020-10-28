package session

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"sync"
)

//定义对象
type MemorySeesionMgr struct {
	sessionMgr map[string]Session
	rowlock    sync.RWMutex
}

//构造函数
func NewMemorySeesionMgr() *MemorySeesionMgr {
	sr := &MemorySeesionMgr{
		sessionMgr: make(map[string]Session, 1024),
	}
	return sr
}

func (s *MemorySeesionMgr) Init(addr string, options ...string) (err error) {
	return
}

func (s *MemorySeesionMgr) CreateSession() (session Session, err error) {
	s.rowlock.Lock()
	defer s.rowlock.Unlock()
	//用UUID作为sessionID
	id := uuid.NewV4()
	//转string
	sessionId := id.String()
	//创建个Session
	session = NewMemorySession(sessionId)
	//加入到大MAP
	s.sessionMgr[sessionId] = session
	return
}

func (s *MemorySeesionMgr) Get(sessionId string) (session Session, err error) {
	s.rowlock.Lock()
	defer s.rowlock.Unlock()
	session, ok := s.sessionMgr[sessionId]
	if ok {
		err = errors.New("session not exists")
		return
	}
	return
}
