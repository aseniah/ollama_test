package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Relative struct {
	FirstName    string `json:"FirstName"`
	LastName     string `json:"LastName"`
	Relationship string `json:"Relationship"`
}

type Person struct {
	FirstName string     `json:"FirstName"`
	LastName  string     `json:"LastName"`
	Birthday  string     `json:"Birthday"`
	Age       int        `json:"Age"`
	Relatives []Relative `json:"Relatives"`
}

func parseDate(dateStr string) (time.Time, error) {
	return time.Parse("1/2/2006", dateStr)
}

func calculateAge(birthDate time.Time, deathDate *time.Time, referenceDate time.Time) int {
	endDate := referenceDate
	if deathDate != nil {
		endDate = *deathDate
	}

	age := endDate.Year() - birthDate.Year()
	// Check if birthday has occurred yet in the year of death/reference
	if endDate.Month() < birthDate.Month() || (endDate.Month() == birthDate.Month() && endDate.Day() < birthDate.Day()) {
		age--
	}
	return age
}

func splitName(fullNama string) (string, string) {
	parts := strings.Fields(fullNama)
	if len(parts) == 0 {
		return "", ""
	}
	if len(parts) == 1 {
		return parts[0], ""
	}
	return parts[0], parts[len(parts)-1]
}

func main() {
	file, err := os.Open("input/input.csv")
	if err != nil {
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return
	}

	if len(records) == 0 {
		fmt.Println("[]")
		return
	}

	// Use the first record as header
	header := records[0]
	data := records[1:]

	referenceDate := time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)
	var people []Person

	for _, row := range data {
		rowMap := make(map[string]string)
		for i, val := range row {
			if i < len(header) {
				rowMap[header[i]] = val
			}
		}

		firstName, lastName := splitName(rowMap["Name"])
		birthT, _ := parseDate(rowMap["Birthday"])
		
		var deathT *time.Time
		if rowMap["Died"] != "" && rowMap["Died"] != "null" {
			t, _ := parseDate(rowMap["Died"])
			deathT = &t
		}

		age := calculateAge(birthT, deathT, referenceDate)

		person := Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthT.Format("2006-01-02"),
			Age:       age,
			Relatives: []Relative{},
		}

		// Process relatives: Father, Mother, Brother, Sister
		relCols := []string{"Father", "Mother", "Brother", "Sister"}
		for _, col := range relCols {
			val := rowMap[col]
			if val != "" && val != "null" {
				rFirstName, rLastName := splitName(val)
				person.Relatives = append(person.Relatives, Relative{
					FirstName:    rFirstName,
					LastName:     rLastName,
					Relationship: col,
				})
			}
		}

		people = append(people, person)
	}

	output, _ := json.MarshalIndent(people, "", "  ")
	fmt.Println(string(output))
}