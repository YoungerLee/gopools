package gopools

import (
	"sync"
)

var (
	pools = make(map[any]*sync.Pool)
	mu    sync.Mutex
)

func Register[K comparable, V any](id K) bool {
	mu.Lock()
	defer mu.Unlock()
	if _, ok := pools[id]; ok { // do not repeatedly register
		return false
	}
	pools[id] = &sync.Pool{New: func() any { return new(V) }}
	return true
}

func Get[K comparable, V any](id K) *V {
	pool, ok := pools[id]
	if !ok {
		return new(V) // new object if pool not exists
	}
	return pool.Get().(*V)
}

func Put[K comparable, V any](id K, v *V) {
	pool, ok := pools[id]
	if !ok {
		return
	}
	pool.Put(v)
}
