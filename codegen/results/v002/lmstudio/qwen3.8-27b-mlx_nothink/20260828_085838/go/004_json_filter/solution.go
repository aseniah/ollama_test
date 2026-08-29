package main

import (
	"encoding/json"
	"fmt"
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

	result := make([]Record, 0)
	for _, r := range records {
		if r.Active && r.Age >= 30 {
			result = append(result, r)
		}
	}

	sort.SliceStable(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})

	out, err := json.Marshal(result)
	if err != nil {
		os.Exit(1)
	}

	fmt.Print(string(out))
}