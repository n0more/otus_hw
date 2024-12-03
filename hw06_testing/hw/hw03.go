package hw

import (
	"fmt"
	"strings"
)

func Runhw03(size int) string {
	output := []string{}

	information := fmt.Sprintf("Размер доски: %v на %v\n", size, size)
	output = append(output, information)

	for i := 0; i < size; i++ {
		for u := 0; u < size; u++ {
			if i%2 == u%2 {
				output = append(output, "#")
			} else {
				output = append(output, "_")
			}
		}
		output = append(output, fmt.Sprintln(""))
	}
	return strings.Join(output, "")
}
