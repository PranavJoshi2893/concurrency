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

// package main

// import (
// 	"fmt"
// 	"runtime"
// 	"time"
// )

// func StartWorker() (stopChan, doneChan chan struct{}) {
// 	stop := make(chan struct{})
// 	done := make(chan struct{})

// 	go func() {
// 		defer close(done)
// 		for {
// 			select {
// 			case <-stop:
// 				return
// 			default:
// 				fmt.Println("woof...")
// 				time.Sleep(time.Millisecond * 500)
// 			}
// 		}
// 	}()

// 	return stop, done
// }

// func main() {
// 	stop1, done1 := StartWorker()
// 	stop2, done2 := StartWorker()

// 	fmt.Println(runtime.NumGoroutine())

// 	time.Sleep(time.Second * 5)
// 	close(stop1)
// 	close(stop2)
// 	<-done1
// 	<-done2

// 	fmt.Println(runtime.NumGoroutine())
// 	fmt.Println("let me SLEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEp... It's 4 a.m")
// }

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func startWorker() (func(), <-chan struct{}) {
	var wg sync.WaitGroup

	stop := make(chan struct{})
	done := make(chan struct{})

	wg.Add(1)

	go func() {
		defer wg.Done()
		defer close(done)
		for {
			select {
			case <-stop:
				return
			case <-time.After(500 * time.Millisecond):
				fmt.Println("running")
			}
		}
	}()

	return func() {
		close(stop)
		wg.Wait()
	}, done
}

func main() {

	stop, done := startWorker()
	fmt.Println(runtime.NumGoroutine())
	time.Sleep(5 * time.Second)
	stop()
	<-done
	fmt.Println(runtime.NumGoroutine())
}
