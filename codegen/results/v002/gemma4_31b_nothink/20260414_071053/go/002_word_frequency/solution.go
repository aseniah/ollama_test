package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

type wordFreq struct {
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
		// Replace non-letters with spaces to avoid merging words
		cleaned := reg.ReplaceAllString(line, " ")
		words := strings.Fields(cleaned)

		for _, w := range words {
			counts[w]++
		}
	}

	// Convert map to slice for sorting
	var freqList []wordFreq
	for w, c := range counts {
		freqList = append(freqList, wordFreq{w, c})
	}

	// Sort: Count descending, then Alphabetical ascending
	sort.Slice(freqList, func(i, j int) bool {
		if freqList[i].count != freqList[j].count {
			return freqList[i].count > freqList[j].count
		}
		return freqList[i].word < freqList[j].word
	})

	// Output
	for _, wf := range freqList {
		fmt.Printf("%s: %d\n", wf.word, wf.count)
	}
}