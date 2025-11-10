// Launch a goroutine that runs forever and figure out how to stop it
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	stop := make(chan struct{})

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-stop:
				return
			default:
				time.Sleep(500 * time.Millisecond)
				fmt.Println("Hello")
			}
		}
	}()

	time.Sleep(4 * time.Second)
	close(stop)
	wg.Wait()

	fmt.Println(runtime.NumGoroutine())

}
