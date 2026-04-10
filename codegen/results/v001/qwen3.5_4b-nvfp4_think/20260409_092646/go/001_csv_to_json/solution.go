package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name   string
	Age    int
	Email  string
	Score  float64
}

func main() {
	file, err := os.Open("input/data.csv")
	if err != nil {
		fmt.Fprintf(os.Stdout, `[]\n`)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	headers, _ := reader.Read([]string{})
	for range headers {
		continue
	}

	var persons []Person

	for {
		row, err := reader.Read()
		if err != nil {
			break
		}

		person := Person{}
		switch len(row) {
		case 4:
			person.Name = row[0]
			person.Age, _ = strconv.Atoi(row[1])
			person.Email = row[2]
			person.Score, _ = strconv.ParseFloat(row[3], 64)
		default:
			continue
		}
		if person.Name == "" {
			continue
		}
		persons = append(persons, person)
	}

	jsonBytes, err := json.Marshal(persons)
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(string(jsonBytes))
}