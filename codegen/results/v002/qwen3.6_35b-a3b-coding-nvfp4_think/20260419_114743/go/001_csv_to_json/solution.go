package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
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
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var people []Person
	for i, record := range records {
		if i == 0 {
			continue // Skip header
		}
		if len(record) < 4 {
			continue
		}

		age, err := strconv.Atoi(record[1])
		if err != nil {
			log.Fatal(err)
		}

		score, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Fatal(err)
		}

		people = append(people, Person{
			Name:  record[0],
			Age:   age,
			Email: record[2],
			Score: score,
		})
	}

	jsonData, err := json.Marshal(people)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(jsonData))
}