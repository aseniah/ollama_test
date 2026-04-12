package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("input/text.txt")
	if err != nil {
		return
	}
	defer file.Close()

	var content string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}

	// Convert to lowercase
	content = strings.ToLower(content)

	// Extract words (only letters)
	re := regexp.MustCompile(`[a-z]+`)
	words := re.FindAllString(content, -1)

	// Count frequencies
	wordCount := make(map[string]int)
	for _, word := range words {
		if word != "" {
			wordCount[word]++
		}
	}

	// Convert map to slice for sorting
	type wordFreq struct {
		word  string
		count int
	}
	var freqs []wordFreq
	for word, count := range wordCount {
		freqs = append(freqs, wordFreq{word, count})
	}

	// Sort by count descending, then alphabetically ascending for ties
	sort.Slice(freqs, func(i, j int) bool {
		if freqs[i].count != freqs[j].count {
			return freqs[i].count > freqs[j].count
		}
		return freqs[i].word < freqs[j].word
	})

	// Output results
	for _, f := range freqs {
		fmt.Printf("%s: %d\n", f.word, f.count)
	}
}