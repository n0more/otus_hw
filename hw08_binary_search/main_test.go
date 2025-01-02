package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hw08(t *testing.T) {
	testCases := []struct {
		name    string
		slice   []int
		tryFind int
		result  bool
	}{
		{
			name:    "true test",
			slice:   []int{5, 90, 11, 34, 2, 3, 10},
			tryFind: 11,
			result:  true,
		},
		{
			name:    "false test",
			slice:   []int{5, 90, 11, 34, 2, 3, 10},
			tryFind: 15,
			result:  false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			testResult := binarySearch(tC.slice, tC.tryFind)
			assert.Equal(t, tC.result, testResult)
		})
	}
}
