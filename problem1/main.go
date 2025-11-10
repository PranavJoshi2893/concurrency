// Launch a goroutine that prints "Hello" 5 times with 100ms delays
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			fmt.Println("Hello")
			time.Sleep(time.Millisecond * 100)
		}
	}()

	wg.Wait()

}
