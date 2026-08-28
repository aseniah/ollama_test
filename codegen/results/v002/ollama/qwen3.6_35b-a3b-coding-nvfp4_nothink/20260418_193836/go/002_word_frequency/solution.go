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

	scanner := bufio.NewScanner(file)
	// Increase buffer size for long lines
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)

	for scanner.Scan() {
		line := scanner.Text()
		// Convert to lowercase
		line = strings.ToLower(line)
		// Extract words: sequences of letters
		re := regexp.MustCompile(`[a-z]+`)
		words := re.FindAllString(line, -1)
		for _, word := range words {
			wordCount[word]++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Collect unique words
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