package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
)

// Record represents the structure of the objects in the input JSON array.
type Record struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Active bool    `json:"active"`
	Score  float64 `json:"score"`
}

func main() {
	// Define the path to the input file
	filePath := "input/data.json"

	// Read the file content
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		// If the file doesn't exist or can't be read, we exit silently 
		// as per the "Do not output anything else" constraint, 
		// but in a real app, you'd handle this error.
		os.Exit(1)
	}

	// Unmarshal the JSON array into a slice of Record
	var allRecords []Record
	if err := json.Unmarshal(data, &allRecords); err != nil {
		os.Exit(1)
	}

	// Filter records: active must be true and age must be >= 30
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

	// Marshal the filtered and sorted slice back to JSON
	output, err := json.Marshal(filteredRecords)
	if err != nil {
		os.Exit(1)
	}

	// Output to stdout
	fmt.Print(string(output))
}