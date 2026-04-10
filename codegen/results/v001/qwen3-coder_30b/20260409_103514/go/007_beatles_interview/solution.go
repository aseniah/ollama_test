package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

type Person struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
}

func main() {
	file, err := os.Open("input/input.csv")
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
	// Skip header row
	for i, record := range records {
		if i == 0 {
			continue
		}
		
		id, err := strconv.Atoi(record[0])
		if err != nil {
			log.Fatal(err)
		}
		
		// Parse birth date (assuming format is MM/DD/YYYY)
		birthDate, err := time.Parse("01/02/2006", record[4])
		if err != nil {
			log.Fatal(err)
		}
		
		// Calculate age as of July 1, 2025
		referenceDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)
		age := referenceDate.Year() - birthDate.Year()
		
		// Adjust if birthday hasn't occurred yet in the reference year
		if referenceDate.Month() < birthDate.Month() || 
		   (referenceDate.Month() == birthDate.Month() && referenceDate.Day() < birthDate.Day()) {
			age--
		}
		
		person := Person{
			ID:        id,
			FirstName: record[1],
			LastName:  record[2],
			Email:     record[3],
			Age:       age,
		}
		
		people = append(people, person)
	}
	
	jsonData, err := json.Marshal(people)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(string(jsonData))
}