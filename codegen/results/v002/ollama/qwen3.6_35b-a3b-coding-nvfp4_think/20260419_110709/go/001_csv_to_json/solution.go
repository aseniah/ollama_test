package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
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
	var people []Person

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		// Skip header row
		if record[0] == "Name" {
			continue
		}

		age, _ := strconv.Atoi(record[1])
		score, _ := strconv.ParseFloat(record[3], 64)

		people = append(people, Person{
			Name:  record[0],
			Age:   age,
			Email: record[2],
			Score: score,
		})
	}

	out, _ := json.Marshal(people)
	fmt.Print(string(out))
}