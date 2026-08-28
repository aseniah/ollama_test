package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

type wordCount struct {
	word  string
	count int
}

func main() {
	file, err := os.Open("input/text.txt")
	if err != nil {
		return
	}
	defer file.Close()

	counts := make(map[string]int)
	// Regex to keep only letters
	reg := regexp.MustCompile(`[^a-z]+`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		// Split by whitespace first to handle word boundaries
		fields := strings.Fields(line)
		for _, f := range fields {
			// Strip all non-letter characters
			cleaned := reg.ReplaceAllString(f, "")
			if cleaned != "" {
				counts[cleaned]++
			}
		}
	}

	// Convert map to slice for sorting
	var sortedList []wordCount
	for w, c := range counts {
		sortedList = append(sortedList, wordCount{w, c})
	}

	// Sort by count descending, then word ascending
	sort.Slice(sortedList, func(i, j int) bool {
		if sortedList[i].count != sortedList[j].count {
			return sortedList[i].count > sortedList[j].count
		}
		return sortedList[i].word < sortedList[j].word
	})

	for _, wc := range sortedList {
		fmt.Printf("%s: %d\n", wc.word, wc.count)
	}
}