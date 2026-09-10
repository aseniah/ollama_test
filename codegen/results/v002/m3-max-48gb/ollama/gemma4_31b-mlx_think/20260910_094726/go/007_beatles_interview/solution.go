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

func splitName(name string) (string, string) {
	parts := strings.Fields(name)
	if len(parts) == 0 {
		return "", ""
	}
	if len(parts) == 1 {
		return parts[0], ""
	}
	return parts[0], parts[len(parts)-1]
}

func calculateAge(birthDate, endDate time.Time) int {
	age := endDate.Year() - birthDate.Year()
	if endDate.Month() < birthDate.Month() || (endDate.Month() == birthDate.Month() && endDate.Day() < birthDate.Day()) {
		age--
	}
	return age
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

	header := records[0]
	data := records[1:]

	var people []Person
	targetDate := time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)

	for _, row := range data {
		rowMap := make(map[string]string)
		for i, val := range row {
			if i < len(header) {
				rowMap[header[i]] = val
			}
		}

		// Parse Name
		firstName, lastName := splitName(rowMap["Name"])

		// Parse Birthday
		birthDate, err := time.Parse("1/2/2006", rowMap["Birthday"])
		if err != nil {
			continue
		}

		// Calculate Age
		endDate := targetDate
		if diedStr, ok := rowMap["Died"]; ok && diedStr != "null" && diedStr != "" {
			if deathDate, err := time.Parse("1/2/2006", diedStr); err == nil {
				endDate = deathDate
			}
		}
		age := calculateAge(birthDate, endDate)

		// Parse Relatives
		var relatives []Relative
		relationships := []string{"Father", "Mother", "Brother", "Sister"}
		for _, relType := range relationships {
			relName := rowMap[relType]
			if relName == "null" || relName == "" {
				continue
			}
			fName, lName := splitName(relName)
			relatives = append(relatives, Relative{
				FirstName:    fName,
				LastName:     lName,
				Relationship: relType,
			})
		}

		people = append(people, Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthDate.Format("2006-01-02"),
			Age:       age,
			Relatives: relatives,
		})
	}

	output, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		return
	}
	fmt.Println(string(output))
}