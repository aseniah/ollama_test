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
		fmt.Fprintf(os.Stderr, "error opening file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading csv: %v\n", err)
		os.Exit(1)
	}

	records := make([]Record, 0, len(rows)-1)
	for i, row := range rows {
		if i == 0 {
			continue // skip header
		}
		age, err := strconv.Atoi(row[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error parsing age: %v\n", err)
			os.Exit(1)
		}
		score, err := strconv.ParseFloat(row[3], 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error parsing score: %v\n", err)
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
		fmt.Fprintf(os.Stderr, "error marshaling json: %v\n", err)
		os.Exit(1)
	}

	fmt.Print(string(output))
}