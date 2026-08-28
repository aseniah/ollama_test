package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

// wordFreq represents a word and its occurrence count.
type wordFreq struct {
	word  string
	count int
}

func main() {
	// The path to the input file as specified.
	const filePath = "input/text.txt"

	// Read the entire content of the file.
	data, err := os.ReadFile(filePath)
	if err != nil {
		// If the file cannot be read, we exit silently to avoid violating
		// the "Do not output anything else" rule.
		return
	}

	// Convert bytes to string and split into tokens by whitespace.
	content := string(data)
	tokens := strings.Fields(content)

	// Use a map to track the frequency of each word.
	counts := make(map[string]int)

	for _, token := range tokens {
		// Convert word to lowercase.
		lowerToken := strings.ToLower(token)

		// Strip all punctuation, keeping only letters.
		var builder strings.Builder
		for _, r := range lowerToken {
			if unicode.IsLetter(r) {
				builder.WriteRune(r)
			}
		}
		cleanWord := builder.String()

		// Only increment if the cleaned word is not an empty string.
		if cleanWord != "" {
			counts[cleanWord]++
		}
	}

	// Convert the map into a slice of wordFreq structs for sorting.
	var freqList []wordFreq
	for word, count := range counts {
		freqList = append(freqList, wordFreq{word: word, count: count})
	}

	// Sort the slice:
	// 1. By count in descending order.
	// 2. For ties, by word in ascending alphabetical order.
	sort.Slice(freqList, func(i, j int) bool {
		if freqList[i].count != freqList[j].count {
			return freqList[i].count > freqList[j].count
		}
		return freqList[i].word < freqList[j].word
	})

	// Output the results in the format "word: count".
	for _, item := range freqList {
		fmt.Printf("%s: %d\n", item.word, item.count)
	}
}