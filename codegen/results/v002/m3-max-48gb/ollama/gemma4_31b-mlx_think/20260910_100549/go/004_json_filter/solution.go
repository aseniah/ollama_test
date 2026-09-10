package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

// Record represents the structure of the data in the JSON file
type Record struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Active bool    `json:"active"`
	Score  float64 `json:"score"`
}

func main() {
	// Read the input file
	filePath := "input/data.json"
	data, err := os.ReadFile(filePath)
	if err != nil {
		return
	}

	var records []Record
	if err := json.Unmarshal(data, &records); err != nil {
		return
	}

	// Filter records: active must be true and age must be 30 or older
	var filtered []Record
	for _, r := range records {
		if r.Active && r.Age >= 30 {
			filtered = append(filtered, r)
		}
	}

	// Sort filtered records by name ascending
	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].Name < filtered[j].Name
	})

	// Convert filtered records back to JSON
	output, err := json.MarshalIndent(filtered, "", "  ")
	if err != nil {
		return
	}

	// Output the result to stdout
	fmt.Println(string(output))
}