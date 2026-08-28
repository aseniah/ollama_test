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

	// Read header
	header, err := reader.Read()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading header: %v\n", err)
		os.Exit(1)
	}

	// Build column index map
	colIndex := make(map[string]int)
	for i, col := range header {
		colIndex[col] = i
	}

	var records []Record

	for {
		row, err := reader.Read()
		if err != nil {
			break // end of file or error
		}

		name := row[colIndex["Name"]]
		ageStr := row[colIndex["Age"]]
		email := row[colIndex["Email"]]
		scoreStr := row[colIndex["Score"]]

		age, err := strconv.Atoi(ageStr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error converting age %s: %v\n", ageStr, err)
			continue
		}

		score, err := strconv.ParseFloat(scoreStr, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error converting score %s: %v\n", scoreStr, err)
			continue
		}

		records = append(records, Record{
			Name:  name,
			Age:   age,
			Email: email,
			Score: score,
		})
	}

	if records == nil {
		records = []Record{}
	}

	output, err := json.Marshal(records)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	os.Stdout.Write(output)
}