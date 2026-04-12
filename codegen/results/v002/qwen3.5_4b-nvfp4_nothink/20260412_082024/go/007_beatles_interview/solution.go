package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Person struct {
	FirstName  string
	LastName   string
	Birthday   string
	Age        int
	Relatives  []Relative
}

type Relative struct {
	FirstName  string
	LastName   string
	Relationship string
}

func main() {
	inputFile := "input/input.csv"
	expectedFormat := "input/expected_format.json"

	// Read input file
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	inputLines := strings.Split(strings.TrimSpace(string(data)), "\n")
	
	// Skip header line and parse rows
	var people []Person
	lines := 1 // start from second line after header
	
	for lines < len(inputLines) && lines <= len(inputLines) {
		rowStr := inputLines[lines]
		if rowStr == "" {
			lines++
			continue
		}
		
		fields := strings.Split(rowStr, ",")
		if len(fields) != 8 {
			panic(fmt.Sprintf("Invalid number of fields at line %d", lines))
		}
		
		name := fields[0]
		birthday := fields[1]
		_ = fields[2] // Died
		
		fatherName := fields[3]
		motherName := fields[4]
		brotherName := fields[5]
		sisterName := fields[6]
		
		// Parse name parts (handle compound names like "John Winston Lennon")
		var firstName, lastName string
		parts := strings.Fields(name)
		if len(parts) > 1 {
			firstName = parts[0]
			lastName = strings.Join(parts[1:], " ")
		} else {
			firstName = name
			lastName = ""
		}
		
		parsedBirth, err := parseDate(birthday)
		if err != nil {
			panic(fmt.Sprintf("Invalid date format at line %d: %v", lines, err))
		}
		
		age, err := calculateAge(parsedBirth, 2025, 7, 1)
		if err != nil {
			panic(err)
		}
		
		var relatives []Relative
		for _, relName := range []string{fatherName, motherName, brotherName, sisterName} {
			if relName != "null" && len(relName) > 0 {
				parts := strings.Fields(relName)
				var relFirst, relLast string
				if len(parts) == 1 {
					relFirst = parts[0]
					relLast = ""
				} else {
					relFirst = parts[0]
					relLast = strings.Join(parts[1:], " ")
				}
				
				var relString string
				switch relName {
				case fields[3]:
					relString = "Father"
				case fields[4]:
					relString = "Mother"
				case fields[5]:
					relString = "Brother"
				case fields[6]:
					relString = "Sister"
				default:
					panic(fmt.Sprintf("Unknown relation type at line %d", lines))
				}
				
				relatives = append(relatives, Relative{
					FirstName:  relFirst,
					LastName:   relLast,
					Relationship: relString,
				})
			}
		}
		
		person := Person{
			FirstName:  firstName,
			LastName:   lastName,
			Birthday:   birthStr(parsedBirth),
			Age:        age,
			Relatives:  relatives,
		}
		people = append(people, person)
		
		lines++
	}
	
	// Generate output JSON
	output, err := json.MarshalIndent(people, "", "   ")
	if err != nil {
		panic(err)
	}
	
	fmt.Print(string(output))
}

func parseDate(date string) (*time.Time, error) {
	var p strings.SplitFunc(func(s string) []string)
	parts := strings.Split(date, "/")
	year, _ := strconv.Atoi(parts[2])
	month, _ := strconv.Atoi(parts[1])
	day, _ := strconv.Atoi(parts[0])
	return time.Date(year, time.Month(month-1), day, 0, 0, 0, 0, time.UTC), nil
}

func calculateAge(from *time.Time, toYear int, toMonth int, toDay int) (int, error) {
	from1970 := from.Year()
	fromMonth := int(from.Month())
	
	resultYear := toYear - from1970
	
	if resultYear < 0 {
		return -resultYear, nil
	}
	
	if fromMonth > toMonth || (fromMonth == toMonth && fromDay > toDay) {
		resultYear--
	}
	
	return resultYear, nil
}

func birthStr(t *time.Time) string {
	return fmt.Sprintf("%02d/%d/%d", t.Day(), t.Month(), t.Year())
}