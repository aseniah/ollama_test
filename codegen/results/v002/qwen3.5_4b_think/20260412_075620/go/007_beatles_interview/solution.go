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
	FirstName  string
	LastName   string
	Relationship string
}

type Person struct {
	FirstName string  `json:"FirstName"`
	LastName  string  `json:"LastName"`
	Birthday  string  `json:"Birthday"`
	Age       int     `json:"Age"`
	Relatives []Relative `json:"Relatives"`
}

func parseName(name string) (firstName, lastName string) {
	parts := strings.Split(name, " ")
	if len(parts) == 1 {
		return parts[0], ""
	}
	
	firstName = parts[0]
	lastName = strings.Join(parts[1:], " ")
	return
}

func parseDate(date string) (day, month, year int, err error) {
	parts := strings.Split(date, "/")
	day, err = strconv.Atoi(parts[0])
	if err != nil {
		return
	}
	
	month, err = strconv.Atoi(parts[1])
	if err != nil {
		return
	}
	
	year, err = strconv.Atoi(parts[2])
	return
}

func calculateAge(birthDate time.Time, refDate time.Time) int {
	years := int(refDate.Year() - birthDate.Year())
	
	if birthDate.Month() > refDate.Month() || 
	   (birthDate.Month() == refDate.Month() && birthDate.Day() >= refDate.Day()) {
		years--
	}
	
	return years
}

func main() {
	file, err := os.Open("input/input.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var people []Person
	refDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)

	for _, record := range records {
		if len(record) < 7 {
			continue
		}
		
		name, err := parseName(record[0])
		if err != nil {
			continue
		}
		
		day, month, year, err := parseDate(record[1])
		if err != nil {
			continue
		}
		
		birthDate := time.Date(year, time.Month(month), day, 0, 0, 0, time.UTC)
		age := calculateAge(birthDate, refDate)
		
		relatives := []Relative{}
		
		father := record[3]
		if father != "null" {
			fatherParts := strings.Split(father, " ")
			fatherName := fatherParts[0]
			if len(fatherParts) > 1 {
				lastPart := ""
				for i := 1; i < len(fatherParts); i++ {
					if lastPart == "" {
						lastPart = fatherParts[i]
					} else {
						lastPart += " " + fatherParts[i]
					}
				}
				relatives = append(relatives, Relative{
					FirstName:  fatherName,
					LastName:   lastPart,
					Relationship: "Father",
				})
			} else {
				relatives = append(relatives, Relative{
					FirstName:  fatherName,
					LastName:   fatherName,
					Relationship: "Father",
				})
			}
		}
		
		mother := record[4]
		if mother != "null" {
			motherParts := strings.Split(mother, " ")
			motherName := motherParts[0]
			if len(motherParts) > 1 {
				lastPart := ""
				for i := 1; i < len(motherParts); i++ {
					if lastPart == "" {
						lastPart = motherParts[i]
					} else {
						lastPart += " " + motherParts[i]
					}
				}
				relatives = append(relatives, Relative{
					FirstName:  motherName,
					LastName:   lastPart,
					Relationship: "Mother",
				})
			} else {
				relatives = append(relatives, Relative{
					FirstName:  motherName,
					LastName:   motherName,
					Relationship: "Mother",
				})
			}
		}
		
		brother := record[5]
		if brother != "null" {
			brotherParts := strings.Split(brother, " ")
			brotherName := brotherParts[0]
			if len(brotherParts) > 1 {
				lastPart := ""
				for i := 1; i < len(brotherParts); i++ {
					if lastPart == "" {
						lastPart = brotherParts[i]
					} else {
						lastPart += " " + brotherParts[i]
					}
				}
				relatives = append(relatives, Relative{
					FirstName:  brotherName,
					LastName:   lastPart,
					Relationship: "Brother",
				})
			} else {
				relatives = append(relatives, Relative{
					FirstName:  brotherName,
					LastName:   brotherName,
					Relationship: "Brother",
				})
			}
		}
		
		sister := record[6]
		if sister != "null" {
			sisterParts := strings.Split(sister, " ")
			sisterName := sisterParts[0]
			if len(sisterParts) > 1 {
				lastPart := ""
				for i := 1; i < len(sisterParts); i++ {
					if lastPart == "" {
						lastPart = sisterParts[i]
					} else {
						lastPart += " " + sisterParts[i]
					}
				}
				relatives = append(relatives, Relative{
					FirstName:  sisterName,
					LastName:   lastPart,
					Relationship: "Sister",
				})
			} else {
				relatives = append(relatives, Relative{
					FirstName:  sisterName,
					LastName:   sisterName,
					Relationship: "Sister",
				})
			}
		}
		
		birthday := fmt.Sprintf("%d-%02d-%02d", year, month, day)
		
		person := Person{
			FirstName: name,
			LastName:  name,
			Birthday:  birthday,
			Age:       age,
			Relatives: relatives,
		}
		people = append(people, person)
	}

	output, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		panic(err)
	}
	
	fmt.Println(string(output))
}