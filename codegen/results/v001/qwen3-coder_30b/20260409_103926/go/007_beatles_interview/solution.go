package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
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
		
		id, _ := strconv.Atoi(record[0])
		firstName := record[1]
		lastName := record[2]
		email := record[3]
		
		// Parse birth date (assuming format is MM/DD/YYYY)
		birthDateStr := record[4]
		birthDate, err := time.Parse("01/02/2006", birthDateStr)
		if err != nil {
			log.Printf("Error parsing date for record %d: %v", i, err)
			continue
		}
		
		// Calculate age as of July 1, 2025
		referenceDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)
		age := referenceDate.Year() - birthDate.Year()
		
		// Adjust if birthday hasn't occurred yet in 2025
		if referenceDate.Month() < birthDate.Month() || 
		   (referenceDate.Month() == birthDate.Month() && referenceDate.Day() < birthDate.Day()) {
			age--
		}
		
		people = append(people, Person{
			ID:        id,
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
			Age:       age,
		})
	}

	output, err := json.Marshal(people)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(output))
}