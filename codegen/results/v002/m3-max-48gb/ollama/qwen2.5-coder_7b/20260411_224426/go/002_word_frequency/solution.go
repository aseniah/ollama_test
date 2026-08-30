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
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	wordCounts := make(map[string]int)

	re := regexp.MustCompile(`[^\w]+`)
	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		words := re.Split(line, -1)
		for _, word := range words {
			if len(word) > 0 {
				wordCounts[word]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	var wcList []wordCount
	for word, count := range wordCounts {
		wcList = append(wcList, wordCount{word: word, count: count})
	}

	sort.SliceStable(wcList, func(i, j int) bool {
		if wcList[i].count == wcList[j].count {
			return wcList[i].word < wcList[j].word
		}
		return wcList[i].count > wcList[j].count
	})

	for _, wc := range wcList {
		fmt.Printf("%s: %d\n", wc.word, wc.count)
	}
}