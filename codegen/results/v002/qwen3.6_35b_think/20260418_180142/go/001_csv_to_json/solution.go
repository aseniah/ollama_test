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
	f, err := os.Open("input/data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var people []Person
	for _, row := range records[1:] {
		if len(row) < 4 {
			continue
		}
		age, err := strconv.Atoi(row[1])
		if err != nil {
			log.Fatal(err)
		}
		score, err := strconv.ParseFloat(row[3], 64)
		if err != nil {
			log.Fatal(err)
		}
		people = append(people, Person{
			Name:  row[0],
			Age:   age,
			Email: row[2],
			Score: score,
		})
	}

	jsonData, err := json.Marshal(people)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonData))
}