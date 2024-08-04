package logger

import "sync"

type TraceNode struct {
	metadata map[string]string
	lock     *sync.RWMutex
}

func NewTraceNode() *TraceNode {
	t := new(TraceNode)
	t.metadata = make(map[string]string, 5)
	t.lock = new(sync.RWMutex)
	return t
}

func (t *TraceNode) Get(key string) string {
	t.lock.RLock()
	defer t.lock.RUnlock()
	return t.metadata[key]
}

func (t *TraceNode) Set(key, val string) {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.metadata[key] = val
}

func (t *TraceNode) ForkMap() map[string]string {
	ret := make(map[string]string, 5)
	t.lock.RLock()
	defer t.lock.RUnlock()
	for k, v := range t.metadata {
		ret[k] = v
	}
	return ret
}

func (t *TraceNode) ForkArbMap() map[string]interface{} {
	ret := make(map[string]interface{})
	t.lock.RLock()
	defer t.lock.RUnlock()
	for k, v := range t.metadata {
		ret[k] = v
	}
	return ret
}
