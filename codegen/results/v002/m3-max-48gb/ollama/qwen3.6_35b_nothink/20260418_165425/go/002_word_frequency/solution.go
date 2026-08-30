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
	re := regexp.MustCompile(`[a-zA-Z]+`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := re.FindAllString(line, -1)
		for _, word := range words {
			lowerWord := strings.ToLower(word)
			wordCount[lowerWord]++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Get all unique words
	words := make([]string, 0, len(wordCount))
	for word := range wordCount {
		words = append(words, word)
	}

	// Sort by count descending, then alphabetically ascending
	sort.Slice(words, func(i, j int) bool {
		countI := wordCount[words[i]]
		countJ := wordCount[words[j]]
		if countI != countJ {
			return countI > countJ
		}
		return words[i] < words[j]
	})

	for _, word := range words {
		fmt.Printf("%s: %d\n", word, wordCount[word])
	}
}