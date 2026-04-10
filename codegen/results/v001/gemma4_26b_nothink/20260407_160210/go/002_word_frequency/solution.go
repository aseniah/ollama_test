package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

// WordFreq stores a word and its associated frequency for sorting
type WordFreq struct {
	Word  string
	Count int
}

func main() {
	// Open the file
	file, err := os.Open("input/text.txt")
	if err != nil {
		// If file doesn't exist or can't be opened, we exit silently per requirements
		// (Do not output anything else)
		return
	}
	defer file.Close()

	wordCounts := make(map[string]int)
	
	// Regex to keep only letters (lowercase handled later)
	// We will strip everything that is not a letter
	re := regexp.MustCompile(`[^a-zA-Z]+`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		
		// Split line into words by whitespace
		words := strings.Fields(line)
		
		for _, w := range words {
			// Convert to lowercase
			w = strings.ToLower(w)
			
			// Strip all punctuation (keep only letters)
			// Replace non-letters with empty string
			cleanWord := re.ReplaceAllString(w, "")
			
			// Only count if the word isn't empty after stripping
			if cleanWord != "" {
				wordCounts[cleanWord]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return
	}

	// Convert map to slice for sorting
	var freqList []WordFreq
	for word, count := range wordCounts {
		freqList = append(freqList, WordFreq{Word: word, Count: count})
	}

	// Sort by count descending; for ties, sort alphabetically ascending
	sort.Slice(freqList, func(i, j int) bool {
		if freqList[i].Count != freqList[j].Count {
			return freqList[i].Count > freqList[j].Count
		}
		return freqList[i].Word < freqList[j].Word
	})

	// Print the results
	for _, item := range freqList {
		fmt.Printf("%s: %d\n", item.Word, item.Count)
	}
}