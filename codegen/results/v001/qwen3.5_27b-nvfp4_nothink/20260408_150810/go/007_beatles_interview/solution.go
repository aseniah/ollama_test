package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type Person struct {
	Name       string `json:"name"`
	BirthDate  string `json:"birth_date"`
	Age        int    `json:"age"`
	Employment string `json:"employment"`
}

func main() {
	file, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening input file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	var people []Person
	var header []string

	// Read and validate header
	if header, err = reader.Read(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading CSV header: %v\n", err)
		os.Exit(1)
	}

	nameIdx := -1
	birthDateIdx := -1
	employmentIdx := -1

	for i, col := range header {
		cleanCol := strings.TrimSpace(col)
		switch cleanCol {
		case "Name":
			nameIdx = i
		case "Birth Date", "Birthdate", "birth_date": // Flexible parsing for common variations
			birthDateIdx = i
		case "Employment", "Status", "employment":
			employmentIdx = i
		}
	}

	if nameIdx == -1 || birthDateIdx == -1 || employmentIdx == -1 {
		fmt.Fprintf(os.Stderr, "Error: Required columns (Name, Birth Date, Employment) not found in header\n")
		os.Exit(1)
	}

	targetDate := time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading CSV row: %v\n", err)
			os.Exit(1)
		}

		// Ensure record has enough columns
		maxIdx := nameIdx
		if birthDateIdx > maxIdx {
			maxIdx = birthDateIdx
		}
		if employmentIdx > maxIdx {
			maxIdx = employmentIdx
		}
		if len(record) <= maxIdx {
			continue // Skip malformed rows if necessary, or error out
		}

		name := strings.TrimSpace(record[nameIdx])
		birthDateStr := strings.TrimSpace(record[birthDateIdx])
		employment := strings.TrimSpace(record[employmentIdx])

		var birthDate time.Time
		// Attempt common date formats found in CSVs
		formats := []string{
			"2006-01-02",
			"01/02/2006",
			"02/01/2006",
			"Jan 2, 2006",
			"January 2, 2006",
		}

		for _, f := range formats {
			if parsed, err := time.Parse(f, birthDateStr); err == nil {
				birthDate = parsed
				break
			}
		}

		var age int
		if !birthDate.IsZero() {
			age = calculateAge(birthDate, targetDate)
		}

		people = append(people, Person{
			Name:       name,
			BirthDate:  birthDateStr, // Keep original string format unless parsing failed
			Age:        age,
			Employment: employment,
		})
	}

	// Output JSON array to stdout
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ") // Optional indentation for readability
	if err := encoder.Encode(people); err != nil {
		fmt.Fprintf(os.Stderr, "Error encoding JSON: %v\n", err)
		os.Exit(1)
	}
}

func calculateAge(birthDate time.Time, referenceDate time.Time) int {
	years := referenceDate.Year() - birthDate.Year()
	if referenceDate.Before(referenceDate.AddDate(-years, 0, 0)) {
		years--
	}
	return years
}