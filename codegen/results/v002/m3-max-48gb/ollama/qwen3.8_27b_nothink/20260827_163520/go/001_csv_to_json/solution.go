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
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading CSV: %v\n", err)
		os.Exit(1)
	}

	if len(records) < 2 {
		fmt.Print("[]\n")
		return
	}

	// Skip header row
	var people []Person
	for i := 1; i < len(records); i++ {
		record := records[i]
		if len(record) < 4 {
			continue
		}
		name := record[0]
		age, err := strconv.Atoi(record[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing age: %v\n", err)
			os.Exit(1)
		}
		email := record[2]
		score, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing score: %v\n", err)
			os.Exit(1)
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