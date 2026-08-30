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

	wordCount := make(map[string]int)
	// Regex to match sequences of letters (a-z)
	re := regexp.MustCompile(`[a-z]+`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		words := re.FindAllString(line, -1)
		for _, word := range words {
			wordCount[word]++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Collect words into a slice for sorting
	type wordCountPair struct {
		word  string
		count int
	}

	pairs := make([]wordCountPair, 0, len(wordCount))
	for word, count := range wordCount {
		pairs = append(pairs, wordCountPair{word, count})
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