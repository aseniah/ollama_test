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
	Name  string `json:"Name"`
	Age   int    `json:"Age"`
	Email string `json:"Email"`
	Score float64 `json:"Score"`
}

func main() {
	file, err := os.Open("input/data.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	_, _ = reader.Read() // Skip header row

	var records []Record

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}

		name := record[0]
		ageStr := record[1]
		email := record[2]
		scoreStr := record[3]

		age, err := strconv.Atoi(ageStr)
		if err != nil {
			continue
		}

		score, err := strconv.ParseFloat(scoreStr, 64)
		if err != nil {
			continue
		}

		r := Record{
			Name:  name,
			Age:   age,
			Email: email,
			Score: score,
		}
		records = append(records, r)
	}

	jsonBytes, err := json.MarshalIndent(records, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		os.Exit(1)
	}

	fmt.Print(string(jsonBytes))
}