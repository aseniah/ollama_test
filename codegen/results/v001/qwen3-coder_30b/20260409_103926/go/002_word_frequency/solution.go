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
			var cleanedWord strings.Builder
			for _, r := range word {
				if unicode.IsLetter(r) {
					cleanedWord.WriteRune(r)
				}
			}
			
			cleaned := cleanedWord.String()
			if cleaned != "" {
				wordCount[cleaned]++
			}
		}
	}
	
	// Convert map to slice for sorting
	type wordFreq struct {
		word  string
		count int
	}
	
	var freqList []wordFreq
	for word, count := range wordCount {
		freqList = append(freqList, wordFreq{word, count})
	}
	
	// Sort by count descending, then by word ascending
	sort.Slice(freqList, func(i, j int) bool {
		if freqList[i].count != freqList[j].count {
			return freqList[i].count > freqList[j].count // descending count
		}
		return freqList[i].word < freqList[j].word // ascending word
	})
	
	// Output results
	for _, wf := range freqList {
		fmt.Printf("%s: %d\n", wf.word, wf.count)
	}
}