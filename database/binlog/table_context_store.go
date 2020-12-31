package binlog

import "sync"

// GlobalContextStore keep track of all table contexts received.
var GlobalContextStore = &tableContextStore{
	cache: map[uint64]TableContext{},
}

type tableContextStore struct {
	mu    sync.Mutex
	cache map[uint64]TableContext
}

func (t *tableContextStore) Set(context TableContext) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.cache[context.TableId()] = context
}

func (t *tableContextStore) Get(id uint64) (context TableContext, ok bool) {
	t.mu.Lock()
	defer t.mu.Unlock()
	context, ok = t.cache[id]
	return
}
