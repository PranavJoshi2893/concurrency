// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func StartWorker() (stopFunc func()) {

// 	var wg sync.WaitGroup

// 	stop := make(chan struct{})

// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		for {
// 			select {
// 			case <-stop:
// 				return
// 			default:
// 				fmt.Println("Working...")
// 				time.Sleep(time.Second)
// 			}
// 		}
// 	}()

// 	return func() {
// 		close(stop)
// 		wg.Wait()
// 	}

// }

// func main() {
// 	stop := StartWorker()
// 	time.Sleep(10 * time.Second)
// 	stop()
// }

package main

import (
	"fmt"
	"time"
	"runtime"
)

func StartWorker() (stopChan, doneChan chan struct{}) {
	stop := make(chan struct{})
	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			select {
			case <-stop:
				return
			default:
				fmt.Println("woof...")
				time.Sleep(time.Millisecond * 500)
			}
		}
	}()

	return stop, done
}

func main() {
	stop1, done1 := StartWorker()
	stop2, done2 := StartWorker()

	fmt.Println(runtime.NumGoroutine())

	time.Sleep(time.Second * 5)
	close(stop1)
	close(stop2)
	<-done1
	<-done2

	fmt.Println(runtime.NumGoroutine())
	fmt.Println("let me SLEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEp... It's 4 a.m")
}
