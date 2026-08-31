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
	f, err := os.Open("input/data.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = 4

	records := make([]Record, 0)

	first := true
	for {
		row, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Fprintf(os.Stderr, "Error reading CSV: %v\n", err)
			os.Exit(1)
		}

		if first {
			first = false
			continue // skip header
		}

		age, err := strconv.Atoi(row[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing Age: %v\n", err)
			os.Exit(1)
		}

		score, err := strconv.ParseFloat(row[3], 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing Score: %v\n", err)
			os.Exit(1)
		}

		records = append(records, Record{
			Name:  row[0],
			Age:   age,
			Email: row[2],
			Score: score,
		})
	}

	output, err := json.Marshal(records)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Print(string(output))
}