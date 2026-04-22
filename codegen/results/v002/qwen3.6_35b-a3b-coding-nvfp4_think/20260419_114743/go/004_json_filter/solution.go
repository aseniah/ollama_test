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
		return
	}

	var records []Record
	if err := json.Unmarshal(data, &records); err != nil {
		return
	}

	var filtered []Record
	for _, r := range records {
		if r.Active && r.Age >= 30 {
			filtered = append(filtered, r)
		}
	}

	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].Name < filtered[j].Name
	})

	out, _ := json.Marshal(filtered)
	fmt.Print(string(out))
}