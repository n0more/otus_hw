package main

import (
	"fmt"
	"sort"
)

func binarySearch(v []int, s int) bool {
	sort.Ints(v)
	low := 0
	high := len(v) - 1

	for low <= high {
		mid := (low + high) / 2
		switch {
		case v[mid] == s:
			return true
		case v[mid] < s:
			low = mid + 1
		default:
			high = mid - 1
		}
	}

	return false
}

func main() {
	s := 15
	v := []int{5, 90, 11, 34, 2, 3, 10}

	if binarySearch(v, s) {
		fmt.Printf("%d найден в cрезе\n", s)
	} else {
		fmt.Printf("%d не найден в cрезе\n", s)
	}
}
