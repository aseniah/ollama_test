package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	// Read the file
	content, err := ioutil.ReadFile("input/data.json")
	if err != nil {
		os.Exit(1)
	}

	var records []Record
	if err := json.Unmarshal(content, &records); err != nil {
		os.Exit(1)
	}

	// Filter records: active == true AND age >= 30
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

	// Convert back to JSON and print to stdout
	output, err := json.Marshal(filtered)
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(string(output))
}