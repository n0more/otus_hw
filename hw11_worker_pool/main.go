package main

import (
	"fmt"
	"sync"
)

func CounterIncreaser(rutines int, counter *int) {
	wg := sync.WaitGroup{}
	var mu sync.Mutex

	for i := 0; i < rutines; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			defer mu.Unlock()

			mu.Lock()

			*counter++
			fmt.Printf("add counter: %d\n", *counter)
		}()
	}

	wg.Wait()
}

func main() {
	var counter int
	CounterIncreaser(100, &counter)
	fmt.Println(counter)
}
