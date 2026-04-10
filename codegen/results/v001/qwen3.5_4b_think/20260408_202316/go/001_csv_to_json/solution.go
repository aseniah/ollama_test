package main

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"strconv"
)

type Person struct {
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Email string  `json:"email"`
	Score float64 `json:"score"`
}

func main() {
	file, err := os.Open("input/data.csv")
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		os.Exit(1)
	}

	var people []Person
	for i, row := range records {
		if i == 0 { // Skip header
			continue
		}
		if len(row) < 4 {
			continue // Skip incomplete rows
		}

		age, err := strconv.Atoi(row[1])
		if err != nil {
			continue
		}

		score, err := strconv.ParseFloat(row[3], 64)
		if err != nil {
			continue
		}

		person := Person{
			Name:  row[0],
			Age:   age,
			Email: row[2],
			Score: score,
		}

		people = append(people, person)
	}

	jsonData, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		os.Exit(1)
	}

	_, err = os.Stdout.Write(jsonData)
	if err != nil {
		os.Exit(1)
	}
}