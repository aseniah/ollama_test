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

func main() {
	// Target date for age calculation
	targetDate := time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)

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

	// Header is: Name, Birthday, Died, Father, Mother, Brother, Sister
	header := records[0]
	var people []Person

	for i := 1; i < len(records); i++ {
		row := records[i]
		rowMap := make(map[string]string)
		for j, val := range row {
			rowMap[header[j]] = val
		}

		nameParts := strings.Split(rowMap["Name"], " ")
		firstName := nameParts[0]
		lastName := nameParts[len(nameParts)-1]

		birthdayStr := rowMap["Birthday"]
		bDay, err := time.Parse("1/2/2006", birthdayStr)
		if err != nil {
			continue
		}

		// Calculate Age
		diedStr := rowMap["Died"]
		var endDate time.Time
		if diedStr == "null" || diedStr == "" {
			endDate = targetDate
		} else {
			dDay, err := time.Parse("1/2/2006", diedStr)
			if err == nil {
				endDate = dDay
			} else {
				endDate = targetDate
			}
		}

		age := endDate.Year() - bDay.Year()
		// Adjust age if the anniversary hasn't occurred yet in the end year
		if endDate.Month() < bDay.Month() || (endDate.Month() == bDay.Month() && endDate.Day() < bDay.Day()) {
			age--
		}

		// Handle Relatives
		var relatives []Relative
		relCols := []string{"Father", "Mother", "Brother", "Sister"}
		for _, relType := range relCols {
			val := rowMap[relType]
			if val == "null" || val == "" {
				continue
			}
			relParts := strings.Split(val, " ")
			relFirst := relParts[0]
			relLast := relParts[len(relParts)-1]
			relatives = append(relatives, Relative{
				FirstName:    relFirst,
				LastName:     relLast,
				Relationship: relType,
			})
		}

		people = append(people, Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  bDay.Format("2006-01-02"),
			Age:       age,
			Relatives: relatives,
		})
	}

	jsonOutput, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		return
	}
	fmt.Println(string(jsonOutput))
}