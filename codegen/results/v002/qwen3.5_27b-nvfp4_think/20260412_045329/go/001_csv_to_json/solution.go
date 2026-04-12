package main

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"strconv"
)

type Person struct {
	Name  string  `json:"Name"`
	Age   int     `json:"Age"`
	Email string  `json:"Email"`
	Score float64 `json:"Score"`
}

func main() {
	file, err := os.Open("input/data.csv")
	if err != nil {
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return
	}

	var people []Person
	for i, record := range records {
		if i == 0 {
			continue // skip header row
		}

		if len(record) < 4 {
			continue
		}

		age, err := strconv.Atoi(record[1])
		if err != nil {
			return
		}

		score, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			return
		}

		person := Person{
			Name:  record[0],
			Age:   age,
			Email: record[2],
			Score: score,
		}
		people = append(people, person)
	}

	jsonOutput, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		return
	}

	os.Stdout.Write(jsonOutput)
}