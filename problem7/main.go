// Create a goroutine that panics and recover from it in the main goroutine
package main

import (
	"fmt"
)

func main() {

	errorCh := make(chan error, 1)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				errorCh <- fmt.Errorf("%v", r)
			}
			close(errorCh)
		}()
		panic("Everything is f*cked up")
	}()

	err := <-errorCh

	if err != nil {
		fmt.Println("Caught error:", err)
	} else {
		fmt.Println("No panic occurred")
	}

}
