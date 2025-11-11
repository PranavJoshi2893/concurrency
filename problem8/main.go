package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"
)

func main() {
	var numGoroutines int

	_, err := fmt.Scan(&numGoroutines)
	if err != nil {
		log.Fatalf("Error at input %v", err)
	}

	if numGoroutines <= 0 {
		log.Fatal("number must me non-negative")
	}

	var wg sync.WaitGroup

	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func(n int) {
			defer wg.Done()
			fmt.Println("Goroutine", n)
		}(i)
	}

	fmt.Println(runtime.NumGoroutine())
	wg.Wait()
	fmt.Println(runtime.NumGoroutine())
}
