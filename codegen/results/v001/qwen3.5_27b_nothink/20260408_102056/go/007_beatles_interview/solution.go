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
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Occupation string `json:"occupation"`
}

func main() {
	// Open the CSV file
	file, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Skip header if present (assuming first line is header based on standard CSV practices)
	_, err = reader.Read()
	if err != nil && err != io.EOF {
		fmt.Fprintf(os.Stderr, "Error reading header: %v\n", err)
		os.Exit(1)
	}

	var people []Person

	// Define the reference date for age calculation: July 1, 2025
	refDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading CSV record: %v\n", err)
			os.Exit(1)
		}

		if len(record) < 3 {
			// Skip malformed lines
			continue
		}

		name := record[0]
		birthDateString := record[1]
		occupation := record[2]

		// Parse birth date. Trying common formats.
		// Assuming format "YYYY-MM-DD" based on typical data, but we'll handle variations.
		var birthDate time.Time
		parsed := false
		
		// Try common formats
		formats := []string{
			"2006-01-02",
			"01/02/2006",
			"02-01-2006",
			"2006/01/02",
			"Jan 2, 2006",
			"January 2, 2006",
		}

		for _, format := range formats {
			if bd, err := time.Parse(format, strings.TrimSpace(birthDateString)); err == nil {
				birthDate = bd
				parsed = true
				break
			}
		}

		if !parsed {
			// If parsing fails, we might default or skip. For this task, we'll skip or assume 0 age.
			// To be robust, we'll create an entry with 0 age or handle gracefully.
			// Given the prompt implies valid transformation, we assume valid dates usually.
			// Let's set age to 0 if unparseable to avoid crashing, though in strict real-world scenarios we might error.
			continue 
		}

		// Calculate age as of July 1, 2025
		age := refDate.Year() - birthDate.Year()
		// Adjust if the birthday hasn't occurred yet in the reference year
		if refDate.YearDay() < birthDate.YearDay() {
			age--
		}
		if refDate.Year() == birthDate.Year() {
			if refDate.YearDay() < birthDate.YearDay() {
				age--
			}
		}
		// Note: Simple logic above handles most cases. 
		// Refined logic:
		// If refDate < birthDate in the same year (not possible if ref is fixed future),
		// Standard age calculation:
		// age = refYear - birthYear
		// if (refMonth, refDay) < (birthMonth, birthDay), age--

		// Let's re-calculate strictly
		age = refDate.Year() - birthDate.Year()
		if (refDate.Month() < birthDate.Month()) || 
		   (refDate.Month() == birthDate.Month() && refDate.Day() < birthDate.Day()) {
			age--
		}

		person := Person{
			Name:       strings.TrimSpace(name),
			Age:        age,
			Occupation: strings.TrimSpace(occupation),
		}

		people = append(people, person)
	}

	// Marshal to JSON
	jsonData, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	// Write to stdout
	fmt.Print(string(jsonData))
}