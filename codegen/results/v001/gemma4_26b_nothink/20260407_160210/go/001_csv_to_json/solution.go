package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Record represents the structure of each row in the CSV
type Record struct {
	Name  string  `json:"Name"`
	Age   int     `json:"Age"`
	Email string  `json:"Email"`
	Score float64 `json:"Score"`
}

func main() {
	// Open the input file
	file, err := os.Open("input/data.csv")
	if err != nil {
		// If file doesn't exist or can't be opened, exit silently or with error
		// Requirement specifies outputting ONLY valid JSON, so we avoid stderr if possible
		// but for a standard utility, os.Exit is common.
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read the header row first to skip it
	_, err = reader.Read()
	if err != nil {
		if err == io.EOF {
			fmt.Print("[]")
			return
		}
		os.Exit(1)
	}

	var records []Record

	// Iterate through the remaining rows
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			os.Exit(1)
		}

		// Ensure we have enough columns to prevent index out of range
		if len(row) < 4 {
			continue
		}

		// Parse Age (int)
		age, err := strconv.Atoi(row[1])
		if err != nil {
			age = 0 // Defaulting or handling error
		}

		// Parse Score (float)
		score, err := strconv.ParseFloat(row[3], 64)
		if err != nil {
			score = 0.0
		}

		// Create the object
		rec := Record{
			Name:  row[0],
			Age:   age,
			Email: row[2],
			Score: score,
		}

		records = append(records, rec)
	}

	// Encode the slice of records to JSON and write to stdout
	encoder := json.NewEncoder(os.Stdout)
	err = encoder.Encode(records)
	if err != nil {
		os.Exit(1)
	}
}