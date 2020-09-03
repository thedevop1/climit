package climit

import (
	"testing"
)

func TestNewLimiter(t *testing.T) {
	l := NewLimiter(0)
	if l.Cap() != 1 {
		t.Errorf("got: %v, want: %v\n", l.Cap(), 1)
	}
}

func TestTryGet(t *testing.T) {
	l := NewLimiter(1)
	if !l.TryGet() {
		t.Errorf("Expect obtain slot, but didn't")
	}
	if l.TryGet() {
		t.Errorf("Expect fail to obtain slot, but got a lot")
	}
}

func Test(t *testing.T) {
	const limit = 5
	l := NewLimiter(limit)
	l.Get()
	for i := 1; i <= limit; i++ {
		if l.TryGet() {
			if i >= limit {
				t.Errorf("Shouldn't get a slot, but did")
			}
		} else {
			if i < limit {
				t.Errorf("Should get a slot, but didn't")
			}
		}
	}
	for i := 0; i < limit; i++ {
		l.Done()
	}
	l.Wait()
}
