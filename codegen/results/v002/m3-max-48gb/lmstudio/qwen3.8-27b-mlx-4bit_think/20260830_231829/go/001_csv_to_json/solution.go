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
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Skip header (row 0), process the rest
	var records []Record
	for i := 1; i < len(rows); i++ {
		row := rows[i]
		if len(row) < 4 {
			continue
		}

		age, err := strconv.Atoi(row[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		score, err := strconv.ParseFloat(row[3], 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
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
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}