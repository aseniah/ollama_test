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

	// Map to store word counts
	wordCounts := make(map[string]int)

	// Read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Convert to lowercase
		line = strings.ToLower(line)
		
		// Use regex to extract only alphabetic words
		re := regexp.MustCompile(`[a-z]+`)
		words := re.FindAllString(line, -1)
		
		for _, word := range words {
			wordCounts[word]++
		}
	}

	// Create a slice of word-count pairs
	type wordCount struct {
		word  string
		count int
	}
	
	var wcList []wordCount
	for word, count := range wordCounts {
		wcList = append(wcList, wordCount{word, count})
	}

	// Sort by count descending, then by word alphabetically ascending
	sort.Slice(wcList, func(i, j int) bool {
		if wcList[i].count != wcList[j].count {
			return wcList[i].count > wcList[j].count
		}
		return wcList[i].word < wcList[j].word
	})

	// Output the results
	for _, wc := range wcList {
		fmt.Printf("%s: %d\n", wc.word, wc.count)
	}
}