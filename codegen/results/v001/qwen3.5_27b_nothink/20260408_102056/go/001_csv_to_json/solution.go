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
		// Only header or empty, output empty array
		jsonBytes, err := json.MarshalIndent([]Person{}, "", "  ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(string(jsonBytes))
		return
	}

	// Skip header row
	var people []Person
	for i := 1; i < len(records); i++ {
		row := records[i]
		if len(row) < 4 {
			// Skip malformed rows
			continue
		}

		name := row[0]
		email := row[2]

		age, err := strconv.Atoi(row[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error converting age for row %d: %v\n", i, err)
			os.Exit(1)
		}

		score, err := strconv.ParseFloat(row[3], 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error converting score for row %d: %v\n", i, err)
			os.Exit(1)
		}

		people = append(people, Person{
			Name:  name,
			Age:   age,
			Email: email,
			Score: score,
		})
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(people); err != nil {
		fmt.Fprintf(os.Stderr, "Error encoding JSON: %v\n", err)
		os.Exit(1)
	}
}