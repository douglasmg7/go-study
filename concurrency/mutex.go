// Run this way to see if there is a race condition.
// go run -race race_condition.go
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	fmt.Println("")

	var counter int
	const gs = 100

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {

			mu.Lock()
			temp := counter
			time.Sleep(time.Millisecond * 1)
			// Free to make some thing else.
			// runtime.Gosched()
			temp++
			counter = temp
			mu.Unlock()

			wg.Done()
		}()
	}

	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("counter:", counter)
}
