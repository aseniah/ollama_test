package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

// wordFreq holds the word and its corresponding frequency for sorting purposes.
type wordFreq struct {
	word  string
	count int
}

func main() {
	// Open the file at the specified path
	filePath := "input/text.txt"
	file, err := os.Open(filePath)
	if err != nil {
		// If the file cannot be opened, we exit silently as per instructions
		return
	}
	defer file.Close()

	counts := make(map[string]int)
	scanner := bufio.NewScanner(file)

	// Process the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		// Split the line by whitespace into tokens
		tokens := strings.Fields(line)

		for _, token := range tokens {
			// Clean the token: strip punctuation and keep only letters, converted to lowercase
			var sb strings.Builder
			for _, r := range token {
				if unicode.IsLetter(r) {
					sb.WriteRune(unicode.ToLower(r))
				}
			}
			
			cleanWord := sb.String()
			// Only add to map if the resulting string is not empty
			if cleanWord != "" {
				counts[cleanWord]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return
	}

	// Convert the map to a slice of structs so we can sort it
	var sortedFreq []wordFreq
	for w, c := range counts {
		sortedFreq = append(sortedFreq, wordFreq{word: w, count: c})
	}

	// Sort the slice:
	// 1. Primary: Frequency descending (count DESC)
	// 2. Secondary: Alphabetical ascending (word ASC)
	sort.Slice(sortedFreq, func(i, j int) bool {
		if sortedFreq[i].count != sortedFreq[j].count {
			return sortedFreq[i].count > sortedFreq[j].count
		}
		return sortedFreq[i].word < sortedFreq[j].word
	})

	// Output the results in the format "word: count"
	for _, item := range sortedFreq {
		fmt.Printf("%s: %d\n", item.word, item.count)
	}
}