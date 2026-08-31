package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Person struct {
	Name  string  `json:"Name"`
	Age   int     `json:"Age"`
	Email string  `json:"Email"`
	Score float64 `json:"Score"`
}

func main() {
	file, err := os.Open("input/data.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	
	// Read all rows
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading CSV: %v\n", err)
		os.Exit(1)
	}

	if len(records) == 0 {
		fmt.Println("[]")
		return
	}

	// Skip header
	var people []Person
	for _, row := range records[1:] {
		if len(row) < 4 {
			continue
		}
		name := row[0]
		age, err1 := strconv.Atoi(row[1])
		email := row[2]
		score, err2 := strconv.ParseFloat(row[3], 64)
		if err1 != nil || err2 != nil {
			continue
		}
		people = append(people, Person{
			Name:  name,
			Age:   age,
			Email: email,
			Score: score,
		})
	}

	output, err := json.Marshal(people)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}