package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
)

type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Active bool   `json:"active"`
	Score  float64 `json:"score"`
}

func main() {
	data, err := ioutil.ReadFile("input/data.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var people []Person
	err = json.Unmarshal(data, &people)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	filteredPeople := make([]Person, 0)
	for _, person := range people {
		if person.Active && person.Age >= 30 {
			filteredPeople = append(filteredPeople, person)
		}
	}

	sort.Slice(filteredPeople, func(i, j int) bool {
		return filteredPeople[i].Name < filteredPeople[j].Name
	})

	jsonData, err := json.MarshalIndent(filteredPeople, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	fmt.Println(string(jsonData))
}