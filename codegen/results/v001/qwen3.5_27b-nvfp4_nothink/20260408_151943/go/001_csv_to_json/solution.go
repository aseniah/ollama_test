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
	
	// Read and skip header
	_, err = reader.Read()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading header: %v\n", err)
		os.Exit(1)
	}

	var records []Record

	for {
		row, err := reader.Read()
		if err != nil {
			break // End of file or error
		}

		// Ensure we have enough columns (index 0: Name, 1: Age, 2: Email, 3: Score)
		if len(row) < 4 {
			fmt.Fprintf(os.Stderr, "Skipping invalid row: %v\n", row)
			continue
		}

		age, err := strconv.Atoi(row[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing age '%s': %v\n", row[1], err)
			os.Exit(1)
		}

		score, err := strconv.ParseFloat(row[3], 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing score '%s': %v\n", row[3], err)
			os.Exit(1)
		}

		rec := Record{
			Name:  row[0],
			Age:   age,
			Email: row[2],
			Score: score,
		}
		records = append(records, rec)
	}

	output, err := json.MarshalIndent(records, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}