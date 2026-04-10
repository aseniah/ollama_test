package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func calculateAge(dob time.Time, referenceDate time.Time) int {
	years := referenceDate.Year() - dob.Year()
	if referenceDate.Month() < dob.Month() ||
		(referenceDate.Month() == dob.Month() && referenceDate.Day() < dob.Day()) {
		years--
	}
	return years
}

func main() {
	referenceDate, _ := time.Parse("2006-01-02", "2025-07-01")

	file, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening CSV file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading CSV: %v\n", err)
		os.Exit(1)
	}

	if len(rows) < 2 {
		fmt.Fprintf(os.Stderr, "CSV file must have header and at least one data row\n")
		os.Exit(1)
	}

	var people []Person
	header := rows[0]
	nameIdx := -1
	dobIdx := -1

	for i, col := range header {
		lowerCol := strings.ToLower(strings.TrimSpace(col))
		if lowerCol == "name" || lowerCol == "full_name" || lowerCol == "first_name" || lowerCol == "last_name" {
			nameIdx = i
		}
		if lowerCol == "dob" || lowerCol == "date_of_birth" || lowerCol == "birth_date" || lowerCol == "birthdate" {
			dobIdx = i
		}
	}

	if nameIdx == -1 || dobIdx == -1 {
		fmt.Fprintf(os.Stderr, "CSV must contain 'name' and 'dob' columns\n")
		os.Exit(1)
	}

	dateFormats := []string{
		"2006-01-02",
		"01/02/2006",
		"02/01/2006",
		"2006/01/02",
		"January 2, 2006",
		"Jan 2, 2006",
		"2 January 2006",
	}

	for i := 1; i < len(rows); i++ {
		row := rows[i]
		if nameIdx >= len(row) || dobIdx >= len(row) {
			continue
		}

		name := strings.TrimSpace(row[nameIdx])
		dobStr := strings.TrimSpace(row[dobIdx])

		var dob time.Time
		for _, format := range dateFormats {
			if parsed, err := time.Parse(format, dobStr); err == nil {
				dob = parsed
				break
			}
		}

		if dob.IsZero() {
			continue
		}

		age := calculateAge(dob, referenceDate)
		if age < 0 || age > 150 {
			continue
		}

		people = append(people, Person{Name: name, Age: age})
	}

	jsonOutput, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(jsonOutput))
}