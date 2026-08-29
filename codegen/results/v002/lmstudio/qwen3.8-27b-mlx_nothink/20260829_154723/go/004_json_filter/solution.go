package main

import (
	"encoding/json"
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
		os.Exit(1)
	}

	var records []Record
	if err := json.Unmarshal(data, &records); err != nil {
		os.Exit(1)
	}

	var result []Record
	for _, r := range records {
		if r.Active && r.Age >= 30 {
			result = append(result, r)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})

	// Ensure we output an array even if empty
	if result == nil {
		result = []Record{}
	}

	out, err := json.Marshal(result)
	if err != nil {
		os.Exit(1)
	}

	os.Stdout.Write(out)
}