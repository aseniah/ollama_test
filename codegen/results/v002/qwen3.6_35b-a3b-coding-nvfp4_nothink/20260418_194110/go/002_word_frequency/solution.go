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
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	wordCounts := make(map[string]int)
	scanner := bufio.NewScanner(file)

	// Regex to match only letters
	re := regexp.MustCompile(`[a-zA-Z]+`)

	for scanner.Scan() {
		line := scanner.Text()
		words := re.FindAllString(line, -1)
		for _, word := range words {
			lowerWord := strings.ToLower(word)
			wordCounts[lowerWord]++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Build a slice of word-count pairs for sorting
	type wordCount struct {
		word  string
		count int
	}

	pairs := make([]wordCount, 0, len(wordCounts))
	for word, count := range wordCounts {
		pairs = append(pairs, wordCount{word, count})
	}

	// Sort by count descending, then by word ascending
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].count != pairs[j].count {
			return pairs[i].count > pairs[j].count
		}
		return pairs[i].word < pairs[j].word
	})

	for _, pair := range pairs {
		fmt.Printf("%s: %d\n", pair.word, pair.count)
	}
}