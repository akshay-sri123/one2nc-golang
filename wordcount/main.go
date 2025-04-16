package main

import (
	"strings"
)

func countLines(text string) int {
	splitLines := strings.Split(text, ". ")
	return len(splitLines)
}

func countWords(text string) int {
	splitWords := strings.Split(text, " ")
	return len(splitWords)
}

func countCharacters(text string) int {
	charCount := 0
	splitWords := strings.Split(text, " ")

	for _, word := range splitWords {
		charCount = charCount + len(word)
	}
	return (charCount + len(splitWords) - 1)
}
