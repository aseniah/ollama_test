package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string   `json:"Name"`
	Age  int      `json:"Age"`
	Email string  `json:"Email"`
	Score float64 `json:"Score"`
}

func main() {
	file, err := os.Open("input/data.csv")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	header, err := reader.Read()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading header:", err)
		os.Exit(1)
	}

	var persons []Person
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading record:", err)
			os.Exit(1)
		}

		person := Person{
			Name:  record[0],
			Age:   0,
			Email: record[2],
			Score: 0.0,
		}

		var age int64
		if _, err = fmt.Sscanf(record[1], "%d", &age); err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing age:", err)
			os.Exit(1)
		}
		person.Age = int(age)

		var score float64
		if _, err = fmt.Sscanf(record[3], "%f", &score); err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing score:", err)
			os.Exit(1)
		}
		person.Score = score

		persons = append(persons, person)
	}

	jsonData, err := json.Marshal(persons)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error marshaling JSON:", err)
		os.Exit(1)
	}

	fmt.Println(string(jsonData))
}