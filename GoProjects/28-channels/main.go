package main

import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println("Channels in Golang")

	waitgroup := &sync.WaitGroup{}
	mychannel := make(chan int, 2)

	waitgroup.Add(2)

	//Receive Only Channel
	go func(wg *sync.WaitGroup, ch <-chan int) {
		fmt.Println("Receiver Goroutine")
		defer wg.Done()
		for value := range ch {
			fmt.Println(value)
		}
	}(waitgroup, mychannel)

	//Send Only Channel
	go func(wg *sync.WaitGroup, ch chan<- int) {
		fmt.Println("Sender Goroutine")
		defer wg.Done()
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}(waitgroup, mychannel)

	// //Can receive and send
	// go func(wg *sync.WaitGroup, ch chan int) {
	// 	defer wg.Done()
	// 	value, isChannelOpen := <-ch
	// 	if isChannelOpen {
	// 		fmt.Println(value)
	// 	}
	// 	ch <- 3
	// 	close(ch)
	// }(waitgroup, mychannel)

	waitgroup.Wait()

}
