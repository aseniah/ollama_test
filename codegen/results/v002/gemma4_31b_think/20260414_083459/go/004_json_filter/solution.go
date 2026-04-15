package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

// Record represents the structure of the input JSON objects.
type Record struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Active bool    `json:"active"`
	Score  float64 `json:"score"`
}

func main() {
	// Read the file input/data.json
	filePath := "input/data.json"
	data, err := os.ReadFile(filePath)
	if err != nil {
		os.Exit(1)
	}

	// Parse the JSON array
	var records []Record
	err = json.Unmarshal(data, &records)
	if err != nil {
		os.Exit(1)
	}

	// Filter records where active is true and age is 30 or older
	var filtered []Record
	for _, r := range records {
		if r.Active && r.Age >= 30 {
			filtered = append(filtered, r)
		}
	}

	// Sort the filtered records by name ascending
	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].Name < filtered[j].Name
	})

	// Marshal the result back to JSON
	output, err := json.MarshalIndent(filtered, "", "  ")
	if err != nil {
		os.Exit(1)
	}

	// Output to stdout
	fmt.Println(string(output))
}