package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	fmt.Println("Before:", runtime.NumGoroutine())

	go func() {
		for {

		}
	}()

	time.Sleep(100 * time.Millisecond)

	count := runtime.NumGoroutine()
	fmt.Println("After:", count)

	if count > 1 {
		fmt.Printf("LEAK DETECTED! %d leaked goroutine(s)\n", count-1)
	}
}
