package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("input/text.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	wordCount := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		for _, word := range words {
			cleaned := cleanWord(word)
			if cleaned != "" {
				wordCount[cleaned]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	type wordFreq struct {
		word  string
		count int
	}

	var freqList []wordFreq
	for word, count := range wordCount {
		freqList = append(freqList, wordFreq{word, count})
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

func cleanWord(word string) string {
	word = strings.ToLower(word)
	var b strings.Builder
	for _, r := range word {
		if 'a' <= r && r <= 'z' {
			b.WriteRune(r)
		}
	}
	return b.String()
}