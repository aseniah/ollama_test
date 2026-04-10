package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Person struct {
	Name string
	BirthDate string
}

func main() {
	// Open the CSV file
	csvFile, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer csvFile.Close()

	// Create a new CSV reader
	reader := csv.NewReader(csvFile)

	// Read all rows from the CSV file
	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Initialize an array to store the Person structs
	var people []Person

	// Loop through each row in the CSV and parse it into a Person struct
	for _, row := range rows[1:] { // Skip the header row
		name := row[0]
		birthDate := row[1]

		people = append(people, Person{
			Name: name,
			BirthDate: birthDate,
		})
	}

	// Define the target date as July 1, 2025
	targetDate := time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)

	// Loop through each person and calculate their age as of the target date
	for i := range people {
		birthDate, err := time.Parse("01/02/2006", people[i].BirthDate)
		if err != nil {
			fmt.Println(err)
			return
		}

		age := targetDate.Year() - birthDate.Year()
		if birthDate.After(targetDate.AddDate(age, 0, 0)) {
			age--
		}

		people[i].Age = age
	}

	// Convert the people slice to JSON and print it to stdout
	jsonData, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonData))
}