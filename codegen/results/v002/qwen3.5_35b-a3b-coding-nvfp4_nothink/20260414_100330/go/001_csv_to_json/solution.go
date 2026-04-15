package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
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
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read header
	header, err := reader.Read()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading header: %v\n", err)
		return
	}

	// Validate header (optional, but good practice)
	expectedHeader := []string{"Name", "Age", "Email", "Score"}
	if len(header) != len(expectedHeader) {
		fmt.Fprintf(os.Stderr, "Unexpected number of columns\n")
		return
	}

	var records []Record

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading row: %v\n", err)
			return
		}

		if len(row) != 4 {
			fmt.Fprintf(os.Stderr, "Unexpected number of columns in row\n")
			continue
		}

		name := row[0]
		ageStr := row[1]
		email := row[2]
		scoreStr := row[3]

		age, err := strconv.Atoi(ageStr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing age: %v\n", err)
			continue
		}

		score, err := strconv.ParseFloat(scoreStr, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing score: %v\n", err)
			continue
		}

		records = append(records, Record{
			Name:  name,
			Age:   age,
			Email: email,
			Score: score,
		})
	}

	jsonOutput, err := json.Marshal(records)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		return
	}

	fmt.Println(string(jsonOutput))
}