package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("input/text.txt")
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	counts := make(map[string]int)
	word := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for _, ch := range line {
			if unicode.IsLetter(ch) {
				word += string(ch)
			} else if word != "" {
				// Word finished (non-letter)
				if len(word) > 0 {
					counts[word]++
					word = ""
				}
			}
		}
		if word != "" {
			counts[word]++
		}
		// Lowercase conversion
		// Note: The above logic accumulates letters without converting to lower.
		// I need to lowercase before adding to map.
		// Let's refactor slightly to keep it clean.
	}
}