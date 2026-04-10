package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

type Record struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Email  string `json:"email"`
}

func main() {
	// Read input CSV
	file, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening CSV: %v\n", err)
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
		fmt.Fprint(os.Stdout, "[]")
		return
	}

	// Skip header
	rows := records[1:]

	// Target date for age calculation
	calcDate := time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)

	var results []Record

	for _, row := range rows {
		if len(row) < 4 {
			continue // Skip malformed rows
		}

		name := strings.TrimSpace(row[0])
		email := strings.TrimSpace(row[2])
		dobStr := strings.TrimSpace(row[3])

		// Parse Date of Birth (assuming YYYY-MM-DD format based on typical CSV data)
		dob, err := time.Parse("2006-01-02", dobStr)
		if err != nil {
			// If parsing fails, we might try other formats or skip. 
			// Given the constraint to infer rules, we assume standard ISO date.
			continue
		}

		// Calculate Age
		age := calcDate.Year() - dob.Year()
		if calcDate.YearDay() < dob.YearDay() {
			age--
		}

		results = append(results, Record{
			Name:  name,
			Age:   age,
			Email: email,
		})
	}

	// Output JSON
	output, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}