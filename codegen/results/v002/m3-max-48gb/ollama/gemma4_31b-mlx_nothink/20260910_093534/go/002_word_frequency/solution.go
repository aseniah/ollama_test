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
	scanner := bufio.NewScanner(file)
	
	// Regex to find sequences of characters (letters only)
	re := regexp.MustCompile(`[a-zA-Z]+`)

	for scanner.Scan() {
		line := scanner.Text()
		// Convert to lowercase and find all word matches
		matches := re.FindAllString(strings.ToLower(line), -1)
		for _, word := range matches {
			counts[word]++
		}
	}

	// Convert map to slice for sorting
	var freqList []wordFreq
	for w, c := range counts {
		freqList = append(freqList, wordFreq{w, c})
	}

	// Sort by count descending, then by word ascending
	sort.Slice(freqList, func(i, j int) bool {
		if freqList[i].count != freqList[j].count {
			return freqList[i].count > freqList[j].count
		}
		return freqList[i].word < freqList[j].word
	})

	// Print output
	for _, wf := range freqList {
		fmt.Printf("%s: %d\n", wf.word, wf.count)
	}
}