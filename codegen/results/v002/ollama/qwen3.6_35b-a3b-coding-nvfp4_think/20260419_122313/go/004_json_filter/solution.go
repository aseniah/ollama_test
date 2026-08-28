package main

import (
	"encoding/json"
	"fmt"
	"log"
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
		log.Fatal(err)
	}

	var records []Record
	if err := json.Unmarshal(data, &records); err != nil {
		log.Fatal(err)
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

	out, err := json.MarshalIndent(filtered, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))
}