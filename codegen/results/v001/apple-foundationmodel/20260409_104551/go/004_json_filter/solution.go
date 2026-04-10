package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Person struct {
	Name string
	Age  int
	Active bool
	Score float64
}

func main() {
	// Read the JSON file
	file, err := ioutil.ReadFile("input/data.json")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Decode JSON into slice of Person structs
	var persons []Person
	if err := json.Unmarshal(file, &persons); err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	// Filter and sort the persons
	filteredAndSorted := persons
		.filter(func(p Person) bool {
			return p.Active && p.Age >= 30
		})
		.sort(func(i, j int) bool {
			return filteredAndSorted[i].Name < filteredAndSorted[j].Name
		})

	// Output the result
	output, err := json.MarshalIndent(filteredAndSorted, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
	}
	fmt.Println(string(output))
}

// Filter function to select persons based on conditions
func (p Persons) filter(f func(Person) bool) Persons {
	return Persons{
		Results: make([]Person, 0, len(p)),
		NumFiltered: 0,
		NumTotal:    len(p),
	}
	for _, p := range p {
		if f(p) {
			p.NumFiltered++
			p.Results = append(p.Results, p)
		}
	}
	return Persons{
		Results: p.Results,
		NumFiltered: p.NumFiltered,
		NumTotal:    p.NumTotal,
	}
}

// Sort function to sort results
func (p Persons) sort(f func(int, int) bool) Persons {
	p.Results = make([]Person, 0, len(p.Results))
	for i, p := range p.Results {
		p.Results = append(p.Results[:i], p.Results[i+1:]...)
	}
	return Persons{
		Results: p.Results,
		NumFiltered: p.NumFiltered,
		NumTotal:    p.NumTotal,
	}
}

// Simple data structure for the slice
type Persons struct {
	Results []Person
	NumFiltered int
	NumTotal  int
}