// channels are the why to talk to the goroutines
// but the inner function working is hidden
// deadlock
// if somebody is listening then only allow to pass the value

package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("channels")

	mych := make(chan int,2)
	wg := &sync.WaitGroup{}

	// mych <- 5
	// fmt.Println(<-mych)

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-mych)
		wg.Done()
	}(mych, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		mych <- 5
		mych <- 6
		close(mych)
		wg.Done()
	}(mych, wg)

	wg.Wait()
}
