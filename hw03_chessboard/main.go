package main

import (
	"fmt"
)

func main() {
	var size int
	// var type bool

	fmt.Printf("Ведите размер доски:")
	fmt.Scanf("%d", &size)

	fmt.Printf("Размер доски: %v на %v\n", size, size)
	fmt.Println("")

	for i := 0; i < size; i++ {
		for u := 0; u < size; u++ {
			if i%2 == u%2 {
				fmt.Printf("#")
			} else {
				fmt.Printf("_")
			}
		}
		fmt.Println("")
	}
}
