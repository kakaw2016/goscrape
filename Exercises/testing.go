package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := 12
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		counter /= 2
		wg.Done()
	}()

	go func() {
		counter *= 0.5
		wg.Done()
	}()

	go func() {
		counter -= 1
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("Final Counter:", max(counter))
}
