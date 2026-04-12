package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

type Person struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Active bool    `json:"active"`
	Score  float64 `json:"score"`
}

func main() {
	data, err := ioutil.ReadFile("input/data.json")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading file:", err)
		os.Exit(1)
	}

	var people []Person
	err = json.Unmarshal(data, &people)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error parsing JSON:", err)
		os.Exit(1)
	}

	var filtered []Person
	for _, p := range people {
		if p.Active && p.Age >= 30 {
			filtered = append(filtered, p)
		}
	}

	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].Name < filtered[j].Name
	})

	output, err := json.MarshalIndent(filtered, "", "   ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error encoding JSON:", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}