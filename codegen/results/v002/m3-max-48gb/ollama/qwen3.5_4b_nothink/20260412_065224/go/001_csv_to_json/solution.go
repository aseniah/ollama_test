package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Person struct {
	Name  string  `json:"Name"`
	Age   int     `json:"Age"`
	Email string  `json:"Email"`
	Score float64 `json:"Score"`
}

func main() {
	// Open the CSV file
	file, err := os.Open("input/data.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read all records
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading CSV: %v\n", err)
		os.Exit(1)
	}

	if len(records) < 1 {
		fmt.Println(`[]`)
		return
	}

	// Skip the header row
	var people []Person

	for i := 1; i < len(records); i++ {
		record := records[i]
		if len(record) != 4 {
			continue
		}

		name := record[0]
		age, err := strconv.Atoi(record[1])
		if err != nil {
			continue
		}
		email := record[2]
		score, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			continue
		}

		people = append(people, Person{
			Name:  name,
			Age:   age,
			Email: email,
			Score: score,
		})
	}

	// Marshal to JSON
	jsonData, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error encoding JSON: %v\n", err)
		os.Exit(1)
	}

	// Print the JSON array to stdout
	fmt.Println(string(jsonData))
}