package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type Record struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Active bool    `json:"active"`
	Score  float64 `json:"score"`
}

func main() {
	data, err := os.ReadFile("input/data.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	var records []Record
	if err := json.Unmarshal(data, &records); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}

	// Filter records where active is true and age is 30 or older
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

	// If no records match, output an empty array
	if len(filtered) == 0 {
		filtered = make([]Record, 0)
	}

	// Output to stdout
	output, err := json.Marshal(filtered)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}