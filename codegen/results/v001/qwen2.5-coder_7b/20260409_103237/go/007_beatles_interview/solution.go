package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Person struct {
	Name    string `json:"name"`
	Birth   string `json:"birth"`
	Address string `json:"address"`
	ID      int    `json:"id"`
}

func calculateAge(birthDate string) (int, error) {
	layout := "2006-01-02"
	bd, err := time.Parse(layout, birthDate)
	if err != nil {
		return 0, err
	}
	now := time.Date(2025, time.July, 1, 0, 0, 0, 0, bd.Location())
	return now.Year() - bd.Year(), nil
}

func main() {
	file, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var people []Person
	for _, record := range records[1:] { // Skip header row
		name := record[0]
		birthDate := record[1]
		address := record[2]
		id, _ := strconv.Atoi(record[3])

		age, err := calculateAge(birthDate)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		person := Person{
			Name:    name,
			Birth:   birthDate,
			Address: address,
			ID:      id,
			Age:     age,
		}
		people = append(people, person)
	}

	jsonOutput, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(jsonOutput))
}