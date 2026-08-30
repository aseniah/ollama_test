package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	file, err := os.Open("input/text.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	wordFreq := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Create regex to keep only letters (a-z), ignoring case and handling punctuation
		pattern := regexp.MustCompile(`[^a-zA-Z]`)
		words := pattern.Split(line, -1)

		for _, word := range words {
			if len(word) > 0 {
				lowerWord := strings.ToLower(word)
				wordFreq[lowerWord]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	var words []string
	for word := range wordFreq {
		words = append(words, word)
	}

	// Sort by count descending, then alphabetically ascending
	sort.Slice(words, func(i, j int) bool {
		if wordFreq[words[i]] != wordFreq[words[j]] {
			return wordFreq[words[i]] > wordFreq[words[j]]
		}
		return words[i] < words[j]
	})

	output := make([]byte, 0)
	for _, word := range words {
		fmt.Printf("%s: %d\n", word, wordFreq[word])
	}
}