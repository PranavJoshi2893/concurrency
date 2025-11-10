// Launch 10 goroutines, each printing its number, and wait for all to complete package main
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	numGorountines := 10

	var wg sync.WaitGroup

	wg.Add(numGorountines)

	for i := 0; i < numGorountines; i++ {
		go func(n int) {
			defer wg.Done()
			fmt.Println(n)
		}(i)
	}

	wg.Wait()
	fmt.Println("currently active:", runtime.NumGoroutine())
}
