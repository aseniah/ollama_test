package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

type Student struct {
	Name  string `json:"Name"`
	Age   int    `json:"Age"`
	Email string `json:"Email"`
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
	
	// Read header to skip it
	_, err = reader.Read()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading header: %v\n", err)
		os.Exit(1)
	}

	var students []Student

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading CSV: %v\n", err)
			os.Exit(1)
		}

		var age int
		var score float64

		// Parse Age as integer
		if _, err := fmt.Sscanf(record[1], "%d", &age); err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing Age: %v\n", err)
			os.Exit(1)
		}

		// Parse Score as float
		if _, err := fmt.Sscanf(record[3], "%f", &score); err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing Score: %v\n", err)
			os.Exit(1)
		}

		students = append(students, Student{
			Name:  record[0],
			Age:   age,
			Email: record[2],
			Score: score,
		})
	}

	// Output as JSON array
	jsonBytes, err := json.MarshalIndent(students, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(jsonBytes))
}