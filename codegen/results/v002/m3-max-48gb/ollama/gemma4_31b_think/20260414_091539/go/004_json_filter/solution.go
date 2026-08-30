package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

// Record defines the structure of the input JSON objects
type Record struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Active bool    `json:"active"`
	Score  float64 `json:"score"`
}

func main() {
	// Read the input file
	data, err := os.ReadFile("input/data.json")
	if err != nil {
		os.Exit(1)
	}

	// Parse JSON array
	var records []Record
	if err := json.Unmarshal(data, &records); err != nil {
		os.Exit(1)
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

	// Marshal result to JSON
	output, err := json.Marshal(filtered)
	if err != nil {
		os.Exit(1)
	}

	// Output to stdout
	fmt.Println(string(output))
}