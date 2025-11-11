package main

import "fmt"

func main() {
	intChan := make(chan int)

	go func() {
		defer close(intChan)
		for i := 0; i < 10; i++ {
			intChan <- i
		}
	}()

	for val := range intChan {
		fmt.Println(val)
	}

}
