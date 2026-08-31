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
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records := make([]Record, 0)

	// Read header
	_, err = reader.Read()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for {
		row, err := reader.Read()
		if err != nil {
			break // End of file or error
		}
		if len(row) < 4 {
			continue
		}

		age, _ := strconv.Atoi(row[1])
		score, _ := strconv.ParseFloat(row[3], 64)

		record := Record{
			Name:  row[0],
			Age:   age,
			Email: row[2],
			Score: score,
		}
		records = append(records, record)
	}

	output, _ := json.MarshalIndent(records, "", "  ")
	fmt.Println(string(output))
}