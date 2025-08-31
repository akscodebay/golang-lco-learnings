package main

import (
	"fmt"
	"sync"
)

var score = []int{0}

func main() {

	mutex := &sync.RWMutex{}
	waitGroup := &sync.WaitGroup{}

	waitGroup.Add(4)

	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		defer wg.Done()
		mut.Lock()
		fmt.Println("func 1 before Score:", score[0])
		score[0]++
		fmt.Println("func 1 after Score:", score[0])
		mut.Unlock()
	}(waitGroup, mutex)

	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		defer wg.Done()
		mut.Lock()
		fmt.Println("func 2 before Score:", score[0])
		score[0]++
		fmt.Println("func 2 after Score:", score[0])
		mut.Unlock()
	}(waitGroup, mutex)

	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		defer wg.Done()
		mut.Lock()
		fmt.Println("func 3 before Score:", score[0])
		score[0]++
		fmt.Println("func 3 after Score:", score[0])
		mut.Unlock()
	}(waitGroup, mutex)

	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		defer wg.Done()
		mut.RLock()
		fmt.Println("func 4 score reader:", score[0])
		mut.RUnlock()
	}(waitGroup, mutex)

	waitGroup.Wait()
}
