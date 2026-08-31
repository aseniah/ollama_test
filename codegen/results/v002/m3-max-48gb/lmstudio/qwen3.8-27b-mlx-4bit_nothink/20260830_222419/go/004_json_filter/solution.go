I'll write a Go program that reads the JSON file, filters records where `active` is true and `age` >= 30, sorts by name ascending, and outputs the result.

<tool_call>
<function=writefile>
<parameter=path>
main.go
</parameter>
<parameter=content>
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type Record struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Active bool    `json:"active"`
	Score  float64 `json:"score"`
}

func main() {
	data, err := os.ReadFile("input/data.json")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var records []Record
	if err := json.Unmarshal(data, &records); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var filtered []Record
	for _, r := range records {
		if r.Active && r.Age >= 30 {
			filtered = append(filtered, r)
		}
	}

	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].Name < filtered[j].Name
	})

	out, err := json.Marshal(filtered)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(string(out))
}

</parameter>
</function>
</tool_call>