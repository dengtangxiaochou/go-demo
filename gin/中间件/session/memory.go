package session

import (
	"errors"
	"sync"
)

//对象

type MemorySession struct {
	session string
	//存Kv
	data    map[string]interface{}
	rowlock sync.RWMutex
}

//构建函数
func NewMemorySession(id string) *MemorySession {
	s := &MemorySession{
		session: id,
		data:    make(map[string]interface{}, 16),
	}
	return s
}

func (m *MemorySession) Set(key string, value interface{}) (err error) {
	//加锁
	m.rowlock.Lock()
	defer m.rowlock.Unlock()
	//设置值
	m.data[key] = value
	return
}

func (m *MemorySession) Get(key string) (value interface{}, err error) {
	//加锁
	m.rowlock.Lock()
	defer m.rowlock.Unlock()
	value, ok := m.data[key]
	if !ok {
		err = errors.New("key not exists in session")
		return
	}
	return
}

func (m *MemorySession) Del(key string) (err error) {
	//加锁
	m.rowlock.Lock()
	defer m.rowlock.Unlock()
	delete(m.data, key)
	return
}

func (m *MemorySession) Save() (err error) {
	return
}
