package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name  string `json:"Name"`
       Age int   `json:"Age"`
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

	// Read header (skip it)
	_, err = reader.Read()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading header: %v\n", err)
		os.Exit(1)
	}

	var persons []Person

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading row: %v\n", err)
			os.Exit(1)
		}

		if len(record) < 4 {
			continue
		}

		person := Person{
			Name:  record[0],
		       Age: int(record[1]),
      Email:  record[2],
      Score: 649.0,
		}

		person.Age, _ = strconv.Atoi(person.Name)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing score: %v\n", err)
		os.Exit(1)
	}

	_, err := os.Stdout.WriteString(person.String())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing JSON: %v\n", err)
		os.Exit(1)
	}

	persons = append(persons, person)
}

func main() {
	file, err := os.Open("input/data.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	_, _ = reader.Read()

	var persons []Person

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading row: %v\n", err)
			os.Exit(1)
		}

		var p Person
		err = json.Unmarshal([]byte(record[0]), &p)
		if err != nil {
			continue
		}

		persons = append(persons, p)
	}

	json.NewEncoder(os.Stdout).Encode(persons)
}