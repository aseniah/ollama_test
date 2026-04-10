package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
)

// Record represents the structure of the objects in the JSON array
type Record struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Active bool    `json:"active"`
	Score  float64 `json:"score"`
}

func main() {
	// Define the input file path
	inputPath := "input/data.json"

	// Read the file content
	data, err := ioutil.ReadFile(inputPath)
	if err != nil {
		// If the file cannot be read, we exit. 
		// In a real-world scenario, you might want to handle this differently,
		// but for this specific requirement, we stop execution.
		os.Exit(1)
	}

	// Unmarshal the JSON array into a slice of Records
	var allRecords []Record
	if err := json.Unmarshal(data, &allRes); err != nil {
		os.Exit(1)
	}

	// Filter the records: active must be true and age must be >= 30
	var filteredRecords []Record
	for _, r := range allRecords {
		if r.Active && r.Age >= 30 {
			filteredRecords = append(filteredRecords, r)
		}
	}

	// Sort the filtered records by Name in ascending order
	sort.Slice(filteredRecords, func(i, j int) bool {
		return filteredRecords[i].Name < filteredRecords[j].Name
	})

	// Encode the result back to JSON and output to stdout
	output, err := json.Marshal(filteredRecords)
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(string(output))
}

// Helper variable to fix the typo in the logic above during implementation
var allRecords []Record