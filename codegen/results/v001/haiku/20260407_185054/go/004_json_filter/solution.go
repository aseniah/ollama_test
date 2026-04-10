package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
)

type Record struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Active bool    `json:"active"`
	Score  float64 `json:"score"`
}

func main() {
	data, err := ioutil.ReadFile("input/data.json")
	if err != nil {
		log.Fatal(err)
	}

	var records []Record
	err = json.Unmarshal(data, &records)
	if err != nil {
		log.Fatal(err)
	}

	var filtered []Record
	for _, record := range records {
		if record.Active && record.Age >= 30 {
			filtered = append(filtered, record)
		}
	}

	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].Name < filtered[j].Name
	})

	output, err := json.Marshal(filtered)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(output))
}
