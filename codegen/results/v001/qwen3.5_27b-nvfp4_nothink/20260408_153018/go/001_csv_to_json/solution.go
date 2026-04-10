package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Record struct {
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
	
	// Read and discard header
	_, err = reader.Read()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading header: %v\n", err)
		os.Exit(1)
	}

	var records []Record

	for {
		row, err := reader.Read()
		if err != nil {
			break // End of file or error (EOF is expected here)
		}

		if len(row) < 4 {
			continue // Skip malformed rows
		}

		name := row[0]
		email := row[2]

		age, err := strconv.Atoi(row[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid age for row: %v\n", row)
			os.Exit(1)
		}

		score, err := strconv.ParseFloat(row[3], 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid score for row: %v\n", row)
			os.Exit(1)
		}

		record := Record{
			Name:  name,
			Age:   age,
			Email: email,
			Score: score,
		}
		records = append(records, record)
	}

	jsonData, err := json.MarshalIndent(records, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(jsonData))
}