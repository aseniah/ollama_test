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
	records := make([]Record, 0)

	// Read header row
	_, err = reader.Read()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading header: %v\n", err)
		os.Exit(1)
	}

	// Read data rows
	for {
		record, err := reader.Read()
		if err != nil {
			break // End of file or error
		}
		if len(record) < 4 {
			continue
		}

		age, err := strconv.Atoi(record[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing age: %v\n", err)
			os.Exit(1)
		}

		score, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing score: %v\n", err)
			os.Exit(1)
		}

		records = append(records, Record{
			Name:  record[0],
			Age:   age,
			Email: record[2],
			Score: score,
		})
	}

	output, err := json.Marshal(records)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}