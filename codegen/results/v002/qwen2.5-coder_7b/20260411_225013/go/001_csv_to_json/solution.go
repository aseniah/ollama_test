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
		person := Person{
			Name:  row[0],
			Age:   parseInt(row[1]),
			Email: row[2],
			Score: parseFloat(row[3]),
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

func parseInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func parseFloat(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}