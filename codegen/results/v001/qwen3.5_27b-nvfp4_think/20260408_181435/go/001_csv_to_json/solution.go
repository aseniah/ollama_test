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
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var people []Person

	for i, record := range records {
		if i == 0 {
			continue // Skip header row
		}
		age, _ := strconv.Atoi(record[1])
		score, _ := strconv.ParseFloat(record[3], 64)

		person := Person{
			Name:  record[0],
			Age:   age,
			Email: record[2],
			Score: score,
		}
		people = append(people, person)
	}

	output, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		panic(err)
	}

	os.Stdout.Write(output)
	os.Stdout.Write([]byte("\n"))
}