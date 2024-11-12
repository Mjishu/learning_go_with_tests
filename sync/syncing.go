package syncing

import "sync"

type Counter struct {
	mu    sync.Mutex
	Calls int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Calls += 1
}

func (c *Counter) Value() int {
	return c.Calls
}

func NewCounter() *Counter {
	return &Counter{}
}
