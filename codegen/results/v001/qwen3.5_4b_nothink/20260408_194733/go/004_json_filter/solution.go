package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
)

type Record struct {
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Active bool   `json:"active"`
	Score float64 `json:"score"`
}

func main() {
	// Read input file
	data, err := ioutil.ReadFile("input/data.json")
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}

	// Parse JSON
	var records []Record
	if err := json.Unmarshal(data, &records); err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	// Filter records: active is true and age is 30 or older
	var filtered []Record
	for _, r := range records {
		if r.Active && r.Age >= 30 {
			filtered = append(filtered, r)
		}
	}

	// Sort by name ascending
	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].Name < filtered[j].Name
	})

	// Output as JSON
	output, err := json.MarshalIndent(filtered, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}

	fmt.Println(string(output))
}