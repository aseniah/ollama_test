package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

// Person represents a person from the CSV input
type Person struct {
	Name string
	Age  int
	BirthDate string
}

func main() {
	// Read input file
	file, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening input file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Parse CSV
	reader := csv.NewReader(file)
	var people []Person

	// Skip header row
	for reader.Read() {
		record := reader.Record
		if len(record) < 2 {
			continue
		}

		name := strings.TrimSpace(record[0])
		age := 0
		if strings.TrimSpace(record[1]) != "" {
			age, _ = strconv.Atoi(strings.TrimSpace(record[1]))
		}
		birthDate := strings.TrimSpace(record[2])

		people = append(people, Person{
			Name:    name,
			Age:     age,
			BirthDate: birthDate,
		})
	}

	// Reference date: July 1, 2025
	referenceDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)

	// Calculate ages based on birth date
	for i := range people {
		if people[i].Age > 0 {
			people[i].Age = referenceDate.Year() - referenceYear(people[i].BirthDate)
		}
	}

	// Convert to JSON format
	jsonOutput, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(jsonOutput))
}