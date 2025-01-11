package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func randomInt(min, max int) int {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max-min+1)))
	if err != nil {
		fmt.Println("Error generating random number:", err)
		return min
	}
	return int(n.Int64()) + min
}

func Generator(duration time.Duration, min int, max int) chan int {
	c := make(chan int)
	go func() {
		timer := time.NewTimer(duration)
		for {
			readInt := randomInt(min, max)
			select {
			case <-timer.C:
				fmt.Println("Generator stopped")
				close(c)
				return
			case c <- readInt:
				fmt.Printf("Generator sent: %d\n", readInt)
				time.Sleep(time.Duration(readInt*10) * time.Millisecond)
			}
		}
	}()
	return c
}

func Accumulator(input chan int, d int) chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		acc := 0
		count := 0
		for v := range input {
			acc += v
			// fmt.Printf("Accumulator received: %d\n", v)
			count++
			if count == d {
				average := acc / d
				// fmt.Printf("Accumulator average: %d\n", average)
				output <- average
				acc = 0
				count = 0
			}
		}
	}()
	return output
}

func main() {
	duration := 1 * time.Minute
	d := 10

	generator := Generator(duration, 1, 50)

	accumulator := Accumulator(generator, d)

	go func() {
		for result := range accumulator {
			fmt.Printf("Reader received average: %d\n", result)
		}
	}()

	time.Sleep(duration + time.Second)
}
