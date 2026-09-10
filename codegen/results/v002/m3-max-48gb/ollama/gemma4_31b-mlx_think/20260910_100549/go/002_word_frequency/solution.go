package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

type wordFreq struct {
	word  string
	count int
}

func main() {
	// Read the file input/text.txt
	content, err := os.ReadFile("input/text.txt")
	if err != nil {
		return
	}

	// Process text: keep only letters, convert to lowercase, and replace others with space
	var cleaned strings.Builder
	for _, r := range string(content) {
		if unicode.IsLetter(r) {
			cleaned.WriteRune(unicode.ToLower(r))
		} else {
			cleaned.WriteRune(' ')
		}
	}

	// Split by whitespace to get individual words
	words := strings.Fields(cleaned.String())

	// Count frequencies
	countsMap := make(map[string]int)
	for _, word := range words {
		countsMap[word]++
	}

	// Convert map to slice for sorting
	var freqList []wordFreq
	for word, count := range countsMap {
		freqList = append(freqList, wordFreq{word, count})
	}

	// Sort by count descending, then by word ascending
	sort.Slice(freqList, func(i, j int) bool {
		if freqList[i].count != freqList[j].count {
			return freqList[i].count > freqList[j].count
		}
		return freqList[i].word < freqList[j].word
	})

	// Output the results in the format "word: count"
	for _, wf := range freqList {
		fmt.Printf("%s: %d\n", wf.word, wf.count)
	}
}