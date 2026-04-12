package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Person struct {
	Name  string `json:"Name"`
	Age   int    `json:"Age"`
	Email string `json:"Email"`
	Score float64 `json:"Score"`
}

func main() {
	file, err := os.Open("input/data.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	var people []Person

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		if len(record) < 4 {
			continue
		}

		name := record[0]
		age, err := strconv.Atoi(record[1])
		if err != nil {
			continue
		}
		email := record[2]
		score, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			continue
		}

		person := Person{
			Name:  name,
			Age:   age,
			Email: email,
			Score: score,
		}

		people = append(people, person)
	}

	output, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}