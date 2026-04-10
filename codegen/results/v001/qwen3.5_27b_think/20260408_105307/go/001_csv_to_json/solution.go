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
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		os.Exit(1)
	}

	if len(records) < 2 {
		os.Stdout.Write([]byte("[]"))
		return
	}

	var people []Person
	for _, record := range records[1:] {
		if len(record) < 4 {
			continue
		}

		name := record[0]

		age, err := strconv.Atoi(record[1])
		if err != nil {
			age = 0
		}

		email := record[2]

		score, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			score = 0.0
		}

		people = append(people, Person{
			Name:  name,
			Age:   age,
			Email: email,
			Score: score,
		})
	}

	output, err := json.Marshal(people)
	if err != nil {
		os.Exit(1)
	}

	os.Stdout.Write(output)
}