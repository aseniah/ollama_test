package main

import (
	"encoding/json"
	"fmt"
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
	// 1. Read the file input/data.json
	content, err := os.ReadFile("input/data.json")
	if err != nil {
		return
	}

	// 2. Parse the JSON array
	var records []Record
	err = json.Unmarshal(content, &records)
	if err != nil {
		return
	}

	// 3. Filter records: active is true and age is 30 or older
	var filtered []Record
	for _, r := range records {
		if r.Active && r.Age >= 30 {
			filtered = append(filtered, r)
		}
	}

	// 4. Sort the filtered records by name ascending
	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].Name < filtered[j].Name
	})

	// 5. Output the resulting JSON array to stdout
	output, err := json.MarshalIndent(filtered, "", "  ")
	if err != nil {
		return
	}

	fmt.Print(string(output))
}