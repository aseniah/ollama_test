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
	file, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading CSV: %v\n", err)
		os.Exit(1)
	}

	// Skip header
	if len(records) < 2 {
		fmt.Println("[]")
		return
	}

	var people []Person
	refDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)

	for _, record := range records[1:] {
		if len(record) < 7 {
			continue
		}
		
		name := strings.TrimSpace(record[0])
		parts := strings.Fields(name)
		var firstName, lastName string
		if len(parts) >= 2 {
			firstName = parts[0]
			lastName = parts[len(parts)-1]
		} else if len(parts) == 1 {
			firstName = parts[0]
			lastName = ""
		}
		
		birthdayStr := strings.TrimSpace(record[1])
		birthday, err := time.Parse("1/2/2006", birthdayStr)
		if err != nil {
			continue
		}
		
		age := refDate.Year() - birthday.Year()
		if refDate.YearDay() < birthday.YearDay() {
			age--
		}
		
		var relatives []Relative
		
		father := strings.TrimSpace(record[3])
		if father != "null" && father != "" {
			fParts := strings.Fields(father)
			if len(fParts) >= 2 {
				relatives = append(relatives, Relative{
					FirstName:    fParts[0],
					LastName:     fParts[len(fParts)-1],
					Relationship: "Father",
				})
			} else if len(fParts) == 1 {
				relatives = append(relatives, Relative{
					FirstName:    fParts[0],
					LastName:     "",
					Relationship: "Father",
				})
			}
		}
		
		mother := strings.TrimSpace(record[4])
		if mother != "null" && mother != "" {
			mParts := strings.Fields(mother)
			if len(mParts) >= 2 {
				relatives = append(relatives, Relative{
					FirstName:    mParts[0],
					LastName:     mParts[len(mParts)-1],
					Relationship: "Mother",
				})
			} else if len(mParts) == 1 {
				relatives = append(relatives, Relative{
					FirstName:    mParts[0],
					LastName:     "",
					Relationship: "Mother",
				})
			}
		}
		
		brother := strings.TrimSpace(record[5])
		if brother != "null" && brother != "" {
			bParts := strings.Fields(brother)
			if len(bParts) >= 2 {
				relatives = append(relatives, Relative{
					FirstName:    bParts[0],
					LastName:     bParts[len(bParts)-1],
					Relationship: "Brother",
				})
			} else if len(bParts) == 1 {
				relatives = append(relatives, Relative{
					FirstName:    bParts[0],
					LastName:     "",
					Relationship: "Brother",
				})
			}
		}
		
		sister := strings.TrimSpace(record[6])
		if sister != "null" && sister != "" {
			sParts := strings.Fields(sister)
			if len(sParts) >= 2 {
				relatives = append(relatives, Relative{
					FirstName:    sParts[0],
					LastName:     sParts[len(sParts)-1],
					Relationship: "Sister",
				})
			} else if len(sParts) == 1 {
				relatives = append(relatives, Relative{
					FirstName:    sParts[0],
					LastName:     "",
					Relationship: "Sister",
				})
			}
		}
		
		if relatives == nil {
			relatives = []Relative{}
		}
		
		people = append(people, Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthday.Format("2006-01-02"),
			Age:       age,
			Relatives: relatives,
		})
	}
	
	if people == nil {
		people = []Person{}
	}
	
	output, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}
	
	fmt.Println(string(output))
}