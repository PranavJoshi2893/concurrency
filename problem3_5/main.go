package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Problem 3 -Create a counter incremented by 1000 goroutines - observe the race condition
// func main() {
// 	var wg sync.WaitGroup
// 	numGoroutines := 1000
// 	counter := 0
// 	wg.Add(numGoroutines)

// 	for i := 0; i < numGoroutines; i++ {
// 		go func() {
// 			defer wg.Done()
// 			counter++
// 		}()
// 	}

// 	wg.Wait()
// 	fmt.Println(counter)
// }

// Problem 4 - Fix problem #3 using sync.Mutex
// func main() {
// 	var wg sync.WaitGroup
// 	var mu sync.Mutex
// 	numGoroutines := 1000
// 	counter := 0
// 	wg.Add(numGoroutines)

// 	for i := 0; i < numGoroutines; i++ {
// 		go func() {
// 			defer wg.Done()

// 			mu.Lock()
// 			defer mu.Unlock()

// 			counter++
// 		}()
// 	}

// 	wg.Wait()
// 	fmt.Println(counter)
// }

// Problem 5 - Fix problem #3 using atomic operations
func main() {
	var wg sync.WaitGroup

	numGoroutines := 1000
	var counter int64 = 0
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}
