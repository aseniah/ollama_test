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
	reg := regexp.MustCompile(`[^a-z ]+`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		// Remove punctuation by replacing non-letters with space or empty
		// To ensure we don't merge words (e.g., "times,it" -> "timesit"), 
		// we replace non-letters with spaces first.
		cleanLine := regexp.MustCompile(`[^a-z\s]+`).ReplaceAllString(line, " ")
		words := strings.Fields(cleanLine)

		for _, w := range words {
			counts[w]++
		}
	}

	var freqList []wordFreq
	for w, c := range counts {
		freqList = append(freqList, wordFreq{w, c})
	}

	sort.Slice(freqList, func(i, j int) bool {
		if freqList[i].count != freqList[j].count {
			return freqList[i].count > freqList[j].count
		}
		return freqList[i].word < freqList[j].word
	})

	for _, wf := range freqList {
		fmt.Printf("%s: %d\n", wf.word, wf.count)
	}
}