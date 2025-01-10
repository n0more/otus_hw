package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Generator(t *testing.T) {
	duration := 10 * time.Second
	test := 100
	generator := Generator(duration, test, test)

	go func() {
		for result := range generator {
			assert.Equal(t, result, test)
		}
	}()

	time.Sleep(duration + time.Second)
}

func Test_Accumulator(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	generator := make(chan int)

	go func() {
		for _, value := range values {
			generator <- value
			fmt.Printf("Generator sent: %d\n", value)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	expected := 5
	d := 10
	accumulator := Accumulator(generator, d)

	go func() {
		for result := range accumulator {
			fmt.Printf("Accumulator result: %d\n", result)
			assert.Equal(t, expected, result)
		}
	}()

	time.Sleep(2 * time.Second)
}

func Test_hw10(t *testing.T) {
	duration := 10 * time.Second
	d := 10
	test := 100
	generator := Generator(duration, test, test)

	accumulator := Accumulator(generator, d)

	go func() {
		for result := range accumulator {
			assert.Equal(t, result, test)
		}
	}()

	time.Sleep(duration + time.Second)
}
