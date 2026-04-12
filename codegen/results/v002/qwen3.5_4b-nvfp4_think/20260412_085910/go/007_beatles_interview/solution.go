package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Relative struct {
	FirstName string `json:"FirstName"`
	LastName   string `json:"LastName"`
	Relationship string `json:"Relationship"`
}

type Person struct {
	FirstName  string `json:"FirstName"`
	LastName   string `json:"LastName"`
	Birthday   string `json:"Birthday"`
	Age        int    `json:"Age"`
	Relatives  []Relative `json:"Relatives"`
}

func main() {
	file, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	var records []Person
	scanner := bufio.Scanner{
		Arr:  true,
	}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		fields := strings.Split(line, ",")
		if len(fields) < 4 {
			continue
		}

		name := fields[0]
		birthdayStr := fields[1]
		hasDied := strings.EqualFold(fields[2], "true")
		fatherName := fields[3]
		motherName := fields[4]
		brotherName := fields[5]
		sisterName := fields[6]

		firstName, lastName := splitName(name)

		record := Person{
			FirstName:  firstName,
			LastName:   lastName,
			Birthday:   birthdayStr,
			Relatives: []Relative{},
		}

		if hasDied {
			record.Relatives = append(record.Relatives, makeRelative(fatherName, "Father"))
			if motherName != "" {
				record.Relatives = append(record.Relatives, makeRelative(motherName, "Mother"))
			}
			record.Relatives = append(record.Relatives, makeRelative(brotherName, "Brother"))
			record.Relatives = append(record.Relatives, makeRelative(sisterName, "Sister"))
		}

		var relatives []PersonRelative
		if hasDied {
			for _, r := range parseFields(fields[3:]) {
				if r.Name != "" && r.Relationship != "" {
					relatives = append(relatives, PersonRelative{
						FirstName:  strings.TrimSpace(r.Name),
						LastName:   strings.TrimSpace(r.FamilyName),
						Relationship: strings.TrimSpace(r.Relationship),
					})
				}
			}
			record.Relatives = append(record.Relatives, relatives...)
		}

		if len(fields) > 6 {
			record.Age = calculateAge(fields[1], 0)
		}

		records = append(records, record)
	}

	output, err := json.MarshalIndent(records, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error marshaling JSON:", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}

func splitName(name string) (first, last string) {
	parts := strings.FieldsFunc(name, func(r rune) bool {
		return r == ' ' || r == '\t'
	})
	if len(parts) > 2 {
		first = strings.TrimSpace(parts[0])
		last = strings.TrimSuffix(strings.Join(parts[1:], " "), ",")
	} else if len(parts) == 2 {
		first = parts[0]
		last = parts[1]
	}
	return first, last
}

func makeRelative(name string, relation string) Relative {
	first, last := splitName(name)
	return Relative{FirstName: first, LastName: last, Relationship: relation}
}

func parseFields(fields []string) []map[string]string {
	result := []map[string]string{}
	for _, field := range fields[3:] {
		fieldsMap := make(map[string]string)
		parts := strings.FieldsFunc(strings.Trim(field, " "), func(r rune) bool {
			return r == ',' || r == ' ' || r == '\t'
		})
		if len(parts) >= 1 {
			fieldsMap["Name"] = strings.TrimSpace(parts[0])
		}
		result = append(result, fieldsMap)
	}
	return result
}

func calculateAge(bdayStr string, referenceYear int) int {
	parts := strings.Split(bdayStr, "/")
	if len(parts) != 3 {
		return 0
	}
	
	var month, day, birthYear int
	for i, p := range parts {
		switch i {
		case 0:
			birthYear, _ = strconv.Atoi(p)
		case 1:
			month, _ = strconv.Atoi(p)
		case 2:
			day, _ = strconv.Atoi(p)
		}
	}

	var age int
	for month <= 6 || (month == 7 && day < 1) {
		age++
	}
	if !hasPassedBirthday(bdayStr, referenceYear) {
		age--
	}
	return age
}

func hasPassedBirthday(bdayStr string, year int) bool {
	parts := strings.Split(bdayStr, "/")
	if len(parts) != 3 {
		return true
	}

	var m, d, y int
	for i, p := range parts {
		switch i {
		case 0:
			y, _ = strconv.Atoi(p)
		case 1:
			m, _ = strconv.Atoi(p)
		case 2:
			d, _ = strconv.Atoi(p)
		}
	}

	return year > y || (year == y && hasPassed(m, d))
}

func hasPassed(m int, d int) bool {
	if m < 7 {
		return true
	}
	if m == 7 && d >= 1 {
		return true
	}
	return false
}

var dateReference = 0