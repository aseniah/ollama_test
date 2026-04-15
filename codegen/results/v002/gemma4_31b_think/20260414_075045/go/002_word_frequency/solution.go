package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

// WordFreq holds the frequency of a specific word
type WordFreq struct {
	Word  string
	Count int
}

func main() {
	// Read the file input/text.txt
	content, err := os.ReadFile("input/text.txt")
	if err != nil {
		// The problem statement implies the file exists; 
		// however, we return silently to avoid extra output if it doesn't.
		return
	}

	// Rule: Convert all words to lowercase
	text := strings.ToLower(string(content))

	// Rule: Strip all punctuation (keep only letters)
	// We replace any non-letter character with a space to ensure 
	// words separated by punctuation are treated as distinct words.
	var cleanedBuilder strings.Builder
	for _, r := range text {
		if unicode.IsLetter(r) {
			cleanedBuilder.WriteRune(r)
		} else {
			cleanedBuilder.WriteRune(' ')
		}
	}

	// Split the cleaned string into individual words based on whitespace
	words := strings.Fields(cleanedBuilder.String())

	// Count the frequency of each word
	counts := make(map[string]int)
	for _, w := range words {
		counts[w]++
	}

	// Move map results into a slice for sorting
	var freqList []WordFreq
	for w, c := range counts {
		freqList = append(freqList, WordFreq{Word: w, Count: c})
	}

	// Rule: Sort by count descending; for ties, sort alphabetically ascending
	sort.Slice(freqList, func(i, j int) bool {
		if freqList[i].Count != freqList[j].Count {
			return freqList[i].Count > freqList[j].Count
		}
		return freqList[i].Word < freqList[j].Word
	})

	// Rule: Output one line per unique word in the format: word: count
	for _, wf := range freqList {
		fmt.Printf("%s: %d\n", wf.Word, wf.Count)
	}
}