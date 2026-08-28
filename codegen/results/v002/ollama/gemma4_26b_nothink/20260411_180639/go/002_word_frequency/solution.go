package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

// WordFreq stores a word and its associated count for sorting
type WordFreq struct {
	Word  string
	Count int
}

func main() {
	// Open the input file
	file, err := os.Open("input/text.txt")
	if err != nil {
		return
	}
	defer file.Close()

	// Map to store word frequencies
	counts := make(map[string]int)

	// Regex to keep only letters (strips punctuation)
	reg := regexp.MustCompile(`[^a-z]+`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Convert line to lowercase
		line := strings.ToLower(scanner.Text())

		// Replace non-letter characters with spaces to separate words correctly
		// e.g., "times," becomes "times "
		cleanedLine := reg.ReplaceAllString(line, " ")

		// Split by whitespace
		words := strings.Fields(cleanedLine)

		for _, w := range words {
			if w != "" {
				counts[w]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return
	}

	// Convert map to a slice of structs for sorting
	var freqList []WordFreq
	for word, count := range counts {
		freqList = append(freqList, WordFreq{Word: word, Count: count})
	}

	// Sort by count descending, then alphabetically ascending
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