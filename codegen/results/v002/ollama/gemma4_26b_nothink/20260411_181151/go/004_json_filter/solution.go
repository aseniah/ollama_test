package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

// Record represents the structure of the input JSON objects
type Record struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Active bool    `json:"active"`
	Score  float64 `json:"score"`
}

func main() {
	// Define the input file path
	const inputPath = "input/data.json"

	// Read the file content
	content, err := ioutil.ReadFile(inputPath)
	if err != nil {
		// If the file doesn't exist or cannot be read, exit silently or handle error
		// Based on instructions, we only output the resulting JSON array.
		os.Exit(1)
	}

	// Unmarshal the JSON array into a slice of Record
	var allRecords []Record
	err = json.Unmarshal(content, &allRes)
	if err != nil {
		os.Exit(1)
	}

	// Filter records: active must be true and age must be 30 or older
	var filteredRecords []Record
	for _, r := range allRecords {
		if r.Active && r.Age >= 30 {
			filteredRecords = append(filteredRecords, r)
		}
	}

	// Sort the filtered records by name ascending
	sort.Slice(filteredRecords, func(i, j int) bool {
		return filteredRecords[i].Name < filteredRecords[j].Name
	})

	// Marshal the result back to JSON
	output, err := json.MarshalIndent(filteredRecords, "", "  ")
	if err != nil {
		os.Exit(1)
	}

	// Output the resulting JSON array to stdout
	fmt.Println(string(output))
}

// Redefining variable scope for the logic above to ensure it compiles as a single block
var allRecords []Record