package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CI(t *testing.T) {
	rutines := 10
	c := 1
	result := 11
	CounterIncreaser(rutines, &c)
	assert.Equal(t, result, c)
}
