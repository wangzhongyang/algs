package main

import (
	"fmt"
	"sync"
)

// 锁机制
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
	}

	go done(&wg)
	fmt.Println("------------------")
	wg.Wait()
	fmt.Println("exit")
}

func add(wg *sync.WaitGroup) {
	wg.Add(1)
}

func done(wg *sync.WaitGroup) {
	for i := 0; i < 100; i++ {
		fmt.Println("down")
		wg.Done()
	}
}
