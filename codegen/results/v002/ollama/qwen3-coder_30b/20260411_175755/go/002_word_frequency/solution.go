package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("input/text.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	wordCount := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		for _, word := range words {
			// Convert to lowercase
			word = strings.ToLower(word)
			// Remove punctuation
			var cleanWord strings.Builder
			for _, r := range word {
				if unicode.IsLetter(r) {
					cleanWord.WriteRune(r)
				}
			}
			cleanWordStr := cleanWord.String()
			if cleanWordStr != "" {
				wordCount[cleanWordStr]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Create a slice of word-count pairs for sorting
	type wordCountPair struct {
		word  string
		count int
	}
	var pairs []wordCountPair
	for word, count := range wordCount {
		pairs = append(pairs, wordCountPair{word, count})
	}

	// Sort by count descending, then by word ascending
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].count != pairs[j].count {
			return pairs[i].count > pairs[j].count // descending count
		}
		return pairs[i].word < pairs[j].word // ascending word
	})

	// Output results
	for _, pair := range pairs {
		fmt.Printf("%s: %d\n", pair.word, pair.count)
	}
}