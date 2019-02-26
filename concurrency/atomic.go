// Run this way to see if there is a race condition.
// go run -race race_condition.go
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("")

	var counter int64
	const gs = 30

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Printf("%v\n ", atomic.LoadInt64(&counter))
			time.Sleep(time.Millisecond * 1)
			// Free to make some thing else.
			// runtime.Gosched()
			wg.Done()
		}()
	}

	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("counter:", counter)
}
