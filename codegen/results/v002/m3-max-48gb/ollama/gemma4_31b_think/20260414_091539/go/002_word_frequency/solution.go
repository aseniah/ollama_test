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
	// Read the input file
	data, err := os.ReadFile("input/text.txt")
	if err != nil {
		return
	}

	content := string(data)

	// Process the text: lowercase and strip punctuation
	// We replace non-letters with spaces to ensure words separated by 
	// punctuation without spaces (e.g., "hello,world") are split correctly.
	var sb strings.Builder
	for _, r := range strings.ToLower(content) {
		if unicode.IsLetter(r) {
			sb.WriteRune(r)
		} else {
			sb.WriteRune(' ')
		}
	}

	// Split by whitespace to get individual words
	words := strings.Fields(sb.String())

	// Count frequencies
	countsMap := make(map[string]int)
	for _, w := range words {
		countsMap[w]++
	}

	// Convert map to slice for sorting
	var freqs []wordFreq
	for w, c := range countsMap {
		freqs = append(freqs, wordFreq{w, c})
	}

	// Sort by count descending, then alphabetically ascending
	sort.Slice(freqs, func(i, j int) bool {
		if freqs[i].count != freqs[j].count {
			return freqs[i].count > freqs[j].count
		}
		return freqs[i].word < freqs[j].word
	})

	// Output results
	for _, wf := range freqs {
		fmt.Printf("%s: %d\n", wf.word, wf.count)
	}
}