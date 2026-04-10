package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	filePath := "input/text.txt"

	// Read file
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading file:", err)
		os.Exit(1)
	}

	text := string(content)

	// Create regex to match words (only letters)
	wordRe := regexp.MustCompile(`[a-z]+`)

	words := wordRe.Split(text, -1)
	for _, w := range words {
		if w != "" {
			text := strings.ToLower(w)
			words = append(words, text)
		}
	}

	// Fix the split - need to match and get the words properly
	var validWords []string
	r := regexp.MustCompile(`[a-zA-Z]+`)
	for _, w := range r.FindAllString(text, -1) {
		validWords = append(validWords, strings.ToLower(w))
	}

	wordCount := make(map[string]int)
	for _, word := range validWords {
		wordCount[word]++
	}

	// Create a slice of struct for sorting
	type WordCount struct {
		Word   string
		Count  int
	}

	var wc []WordCount
	for word, count := range wordCount {
		wc = append(wc, WordCount{word, count})
	}

	// Sort: descending by count, then ascending alphabetically for ties
	sort.Slice(wc, func(i, j int) bool {
		if wc[i].Count != wc[j].Count {
			return wc[i].Count > wc[j].Count
		}
		return wc[i].Word < wc[j].Word
	})

	// Output results
	for _, wc := range wc {
		fmt.Printf("%s: %d\n", wc.Word, wc.Count)
	}
}