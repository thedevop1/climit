package climit_test

import (
	"fmt"
	"time"

	"github.com/thedevop1/climit"
)

type DB struct {
	limiter *climit.Limiter
}

func (d *DB) Call() {
	defer d.limiter.Done()
	// ...
}

func Example() {
	db := DB{limiter: climit.NewLimiter(5)}
	for i := 0; i < 10; i++ {
		db.limiter.Get() // Block until an available slot
		go db.Call()
	}
	db.limiter.Wait()
}

func ExampleLimiter_TryGet() {
	l := climit.NewLimiter(1)
	for i := 0; i < 3; i++ {
		if l.TryGet() { // Try non-blocking obtain a slot
			// Obtained a slot
			go func(i int) {
				defer l.Done()
				time.Sleep(100 * time.Millisecond)
				fmt.Println(i, "obtained a slot")
			}(i)
		} else {
			// All available slots in use
			fmt.Println(i, "didn't get a slot")
		}
	}
	l.Wait()
	// Output:
	// 1 didn't get a slot
	// 2 didn't get a slot
	// 0 obtained a slot
}
