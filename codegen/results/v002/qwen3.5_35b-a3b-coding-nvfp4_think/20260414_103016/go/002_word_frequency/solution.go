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
	// Read the file
	file, err := os.Open("input/text.txt")
	if err != nil {
		return
	}
	defer file.Close()

	// Count word frequencies
	wordCounts := make(map[string]int)
	scanner := bufio.NewScanner(file)

	// Regex to match only letters
	regex := regexp.MustCompile(`[a-zA-Z]+`)

	for scanner.Scan() {
		line := scanner.Text()
		matches := regex.FindAllString(line, -1)
		for _, match := range matches {
			word := strings.ToLower(match)
			wordCounts[word]++
		}
	}

	if err := scanner.Err(); err != nil {
		return
	}

	// Create a slice of word-count pairs
	type wordCount struct {
		word  string
		count int
	}

	var words []wordCount
	for word, count := range wordCounts {
		words = append(words, wordCount{word: word, count: count})
	}

	// Sort by count descending, then alphabetically ascending
	sort.Slice(words, func(i, j int) bool {
		if words[i].count != words[j].count {
			return words[i].count > words[j].count
		}
		return words[i].word < words[j].word
	})

	// Output results
	for _, wc := range words {
		fmt.Printf("%s: %d\n", wc.word, wc.count)
	}
}