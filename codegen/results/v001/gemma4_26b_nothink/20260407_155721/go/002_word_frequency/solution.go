package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

// wordFreq holds a word and its frequency for sorting purposes
type wordFreq struct {
	word  string
	count int
}

func main() {
	// Open the file
	file, err := os.Open("input/text.txt")
	if err != nil {
		return
	}
	defer file.Close()

	counts := make(map[string]int)
	
	// Regex to keep only letters
	reg, err := regexp.Compile("[^a-z]+")
	if err != nil {
		return
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Convert to lowercase
		line := strings.ToLower(scanner.Text())
		
		// Split by whitespace
		words := strings.Fields(line)
		
		for _, w := range words {
			// Strip punctuation (keep only letters)
			cleanWord := reg.ReplaceAllString(w, "")
			
			if cleanWord != "" {
				counts[cleanWord]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return
	}

	// Convert map to slice for sorting
	var sortedFreqs []wordFreq
	for word, count := range counts {
		sortedFreqs = append(sortedFreqs, wordFreq{word, count})
	}

	// Sort by count descending, then alphabetically ascending
	sort.Slice(sortedFreqs, func(i, j int) bool {
		if sortedFreqs[i].count != sortedFreqs[j].count {
			return sortedFreqs[i].count > sortedFreqs[j].count
		}
		return sortedFreqs[i].word < sortedFreqs[j].word
	})

	// Output the results
	for _, wf := range sortedFreqs {
		fmt.Printf("%s: %d\n", wf.word, wf.count)
	}
}