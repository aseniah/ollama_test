package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

// Record defines the structure of the JSON objects in the input file
type Record struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Active bool    `json:"active"`
	Score  float64 `json:"score"`
}

func main() {
	// Open the input file
	file, err := os.Open("input/data.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Decode the JSON array from the file
	var records []Record
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&records); err != nil {
		fmt.Fprintf(os.Stderr, "Error decoding JSON: %v\n", err)
		os.Exit(1)
	}

	// Filter records: active must be true and age must be 30 or older
	var filtered []Record
	for _, r := range records {
		if r.Active && r.Age >= 30 {
			filtered = append(filtered, r)
		}
	}

	// Sort records by name ascending
	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].Name < filtered[j].Name
	})

	// Encode the resulting slice back to JSON and output to stdout
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ") // Set indent for readability as per common JSON outputs
	if err := encoder.Encode(filtered); err != nil {
		fmt.Fprintf(os.Stderr, "Error encoding JSON: %v\n", err)
		os.Exit(1)
	}
}