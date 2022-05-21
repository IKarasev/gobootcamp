package main

import (
	"fmt"
	"sync"
)

func deposit(b *int, n int, wg *sync.WaitGroup, mute *sync.Mutex) {
	mute.Lock()
	defer mute.Unlock()

	*b += n
	wg.Done()
}

func withdraw(b *int, n int, wg *sync.WaitGroup, mute *sync.Mutex) {
	mute.Lock()
	defer mute.Unlock()

	*b -= n
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var mute sync.Mutex

	wg.Add(200)

	balance := 100

	for i := 0; i < 100; i++ {
		go deposit(&balance, i, &wg, &mute)
		go withdraw(&balance, i, &wg, &mute)
	}
	wg.Wait()

	fmt.Println("Final balance value:", balance)
}
