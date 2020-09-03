// Package climit provides simple functions to limit concurrency. It has similar API as Waitgroup.
package climit

// Limiter limits concurrency. Use Get to obtain a slot, and Done when finished. Wait can be used
// to block until all slots are free.
type Limiter struct {
	ch chan bool
}

// NewLimiter creates a new limiter with capacity of limit.
func NewLimiter(limit int) *Limiter {
	if limit < 1 {
		limit = 1
	}
	return &Limiter{make(chan bool, limit)}
}

// Cap returns the capacity of the limiter.
func (l *Limiter) Cap() int {
	return cap(l.ch)
}

// Get obtains 1 slot, blocks until a slot is available.
func (l *Limiter) Get() {
	l.ch <- true
}

// TryGet uses nonblocking method to obtains 1 slot, returns true if obtained and false otherwise.
func (l *Limiter) TryGet() bool {
	select {
	case l.ch <- true:
		return true
	default:
		return false
	}
}

// Done returns 1 slot.
func (l *Limiter) Done() {
	select {
	case <-l.ch:
		return
	default:
		panic("climit: negative slots")
	}
}

// Wait blocks until all the slots returned. Do not call concurrently with Get or TryGet.
func (l *Limiter) Wait() {
	for i := 0; i < cap(l.ch); i++ {
		l.ch <- true
	}
	l.ch = make(chan bool, l.Cap())
}
