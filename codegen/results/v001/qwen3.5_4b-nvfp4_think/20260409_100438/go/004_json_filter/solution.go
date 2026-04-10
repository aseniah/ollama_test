package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type Record struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Active bool   `json:"active"`
	Score float64 `json:"score"`
}

func main() {
	data, err := os.ReadFile("input/data.json")
	if err != nil {
		os.Exit(1)
	}

	var records []Record
	err = json.Unmarshal(data, &records)
	if err != nil {
		os.Exit(1)
	}

	filtered := make([]Record, 0, len(records))
	for _, r := range records {
		if r.Active && r.Age >= 30 {
			filtered = append(filtered, r)
		}
	}

	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].Name < filtered[j].Name
	})

	output, err := json.Marshal(filtered)
	if err != nil {
		os.Exit(1)
	}
	fmt.Print(string(output))
}