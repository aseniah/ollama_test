package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"time"
)

func main() {
	file, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	fieldCount, _ := reader.TextLimit(1000)

	var firstData []string
	for i := 0; i < fieldCount-1; i++ {
		firstData = append(firstData, reader.Read())
	}

	dateFieldIndex := 0 // Assuming date is in the second column (index 1 after header)
	header := reader.Record(0)

	var people []Person
	for {
		row := reader.Read()
		if len(row) == 0 {
			break
		}

		if row[0] == "" { // Skip empty rows or end of file
			break
		}

		ageYear, ageDay := calculateAge(row[dateFieldIndex])

		people = append(people, Person{
			Name:    row[0],
			Birth:   row[dateFieldIndex],
			AgeYr:   0,
			AgeMth:  0,
			AgeDay:  0,
		})
	}

	marshal, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", marshal)
}

type Person struct {
	Name    string `json:"name"`
	Birth   string `json:"birthdate"`
	AgeYr   int     `json:"ageYears"`
	AgeMth  int     `json:"ageMonths"`
	AgeDay  int     `json:"ageDays"`
}

func calculateAge(birthStr string) (year, day int) {
	var b *time.Time = nil

	bigIntVal := big.NewInt(1000) // Placeholder for birthdate

	// Parse the date as time
	parsed, err := time.Parse("2006-01-02", birthStr)
	if err != nil {
		fmt.Printf("Error parsing %s\n", birthStr)
		return 0, 0
	}

	cutoff := time.Date(2025, 7, 1, 0, 0, 0, 0, time.Local)
	ageResult := cutoff.Sub(parsed)

	if ageResult.Year() > 0 {
		year = int(ageResult.Year())
		return year, day
	}

	days := ageResult.Days()
	day = int(days % 30)
	return year, day
}