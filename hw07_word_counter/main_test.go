package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hw07(t *testing.T) {
	testCases := []struct {
		name  string
		text  string
		word  string
		count int
	}{
		{
			name:  "first test",
			text:  "test, test, one test and! no more test",
			word:  "test",
			count: 4,
		},
		{
			name:  "last test",
			text:  "test, test, one test and! no more test",
			word:  "and",
			count: 1,
		},
		{
			name:  "hello test",
			text:  "Hello, world! или Привет мир. Please, мир  Hello, hello World",
			word:  "hello",
			count: 3,
		},
		{
			name:  "hello test2",
			text:  "Hello, world! или Привет мир. Please, мир  Hello, hello World",
			word:  "world",
			count: 2,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			got := countWords(tC.text)
			countWord := got[tC.word]
			assert.Equal(t, tC.count, countWord)
		})
	}
}
