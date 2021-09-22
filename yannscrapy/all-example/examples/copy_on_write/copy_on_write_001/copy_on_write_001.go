package main

import (
	"sync"
	"sync/atomic"
)

func main() {
	type Map map[string]string
	var m atomic.Value
	m.Store(make(Map))
	m.Store(make(Map))
	mutex := sync.Mutex{}

	read := func(key string) string {
		m1 := m.Load().(Map)
		return m1[key]
	}

	insert := func(key, val string) {
		mutex.Lock()
		defer mutex.Unlock()
		m1 := m.Load().(Map)
		newM := make(Map)
		for k, v := range m1 {
			newM[k] = v
		}
		newM[key] = val
		m.Store(newM)
	}


	_, _ = read, insert
}























