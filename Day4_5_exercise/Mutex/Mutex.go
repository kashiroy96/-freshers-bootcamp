package mutex

import "sync"

type mutex struct {
	m sync.Map
}

var Mutex mutex

func (m *mutex) Lock(key string) bool {
	_, ok := m.m.Load(key)
	if ok {
		return false
	}

	m.m.Store(key, true)
	return !ok
}

func (m *mutex) UnLock(id string) {
	m.m.Delete(id)
}
