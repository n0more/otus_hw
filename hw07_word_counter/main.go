package main

import (
	"strings"
	"unicode"
)

func countWords(s string) map[string]int {
	words := strings.Fields(s)
	wordCount := make(map[string]int)
	for _, word := range words {
		if unicode.IsLetter(rune(word[0])) {
			word = strings.ToLower(word)
			cleanWord := strings.TrimFunc(word, func(r rune) bool {
				return !unicode.IsLetter(r) && !unicode.IsNumber(r)
			})
			wordCount[cleanWord]++
		}
	}
	return wordCount
}

func main() {
	text := "Hello, world! или Привет мир. Please, мир  Hello, hello World"
	textWordCount := countWords(text)
	for word, count := range textWordCount {
		println(word, ":", count)
	}
}
