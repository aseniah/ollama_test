package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

// wordFreq stores a word and its occurrence count
type wordFreq struct {
	word  string
	count int
}

func main() {
	// Read the input file
	content, err := os.ReadFile("input/text.txt")
	if err != nil {
		// If the file cannot be read, we exit silently as per requirements
		// to avoid outputting anything else.
		return
	}

	// Map to store the frequency of each word
	counts := make(map[string]int)

	// Split content into words by whitespace
	words := strings.Fields(string(content))

	for _, w := range words {
		// Strip punctuation and keep only letters
		var sb strings.Builder
		for _, r := range w {
			if unicode.IsLetter(r) {
				sb.WriteRune(unicode.ToLower(r))
			}
		}
		
		cleaned := sb.String()
		
		// Only count if the word contains letters after cleaning
		if cleaned != "" {
			counts[cleaned]++
		}
	}

	// Convert map to a slice of structs for sorting
	var freqList []wordFreq
	for word, count := range counts {
		freqList = append(freqList, wordFreq{word, count})
	}

	// Sort the slice:
	// 1. By count descending
	// 2. By word ascending (alphabetical) for ties
	sort.Slice(freqList, func(i, j int) bool {
		if freqList[i].count != freqList[j].count {
			return freqList[i].count > freqList[j].count
		}
		return freqList[i].word < freqList[j].word
	})

	// Output the results
	for _, item := range freqList {
		fmt.Printf("%s: %d\n", item.word, item.count)
	}
}