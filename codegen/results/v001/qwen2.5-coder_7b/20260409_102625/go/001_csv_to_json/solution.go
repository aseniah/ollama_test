package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
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
		fmt.Println(err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	var people []Person

	for i, row := range rows {
		if i == 0 {
			continue // Skip the header row
		}

		name := row[0]
		age, _ := strconv.Atoi(row[1])
		email := row[2]
		score, _ := strconv.ParseFloat(row[3], 64)

		person := Person{
			Name:  name,
			Age:   age,
			Email: email,
			Score: score,
		}

		people = append(people, person)
	}

	jsonData, err := json.MarshalIndent(people, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonData))
}