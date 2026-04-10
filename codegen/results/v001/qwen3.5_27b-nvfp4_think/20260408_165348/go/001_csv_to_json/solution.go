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
	rows, err := reader.ReadAll()
	if err != nil {
		os.Exit(1)
	}

	var people []Person
	for i := 1; i < len(rows); i++ {
		row := rows[i]
		if len(row) < 4 {
			continue
		}

		age, _ := strconv.Atoi(row[1])
		score, _ := strconv.ParseFloat(row[3], 64)

		people = append(people, Person{
			Name:  row[0],
			Age:   age,
			Email: row[2],
			Score: score,
		})
	}

	encoder := json.NewEncoder(os.Stdout)
	if err := encoder.Encode(people); err != nil {
		os.Exit(1)
	}
}