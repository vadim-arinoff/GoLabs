package main

import (
	"fmt"
	"strings"
)

// findLongestWord находит самое длинное слово в строке.
func findLongestWord(text string) string {
	words := strings.Fields(text) // Разбиваем строку на слова.

	longestWord := ""
	for _, word := range words {
		if len(word) > len(longestWord) {
			longestWord = word
		}
	}

	return longestWord
}

func main() {
	text := "Пример текста с различными словами"
	longest := findLongestWord(text)
	fmt.Printf("Самое длинное слово: %s\n", longest)
}
