package main

import (
    "encoding/json"
    "os"
)

type Record struct {
    Name  string  `json:"name"`
    Age   int     `json:"age"`
    Active bool    `json:"active"`
    Score float64 `json:"score"`
}

func main() {
    file, err := os.Open("input/data.json")
    if err != nil {
        return
    }
    defer file.Close()

    var data []Record
    decoder := json.NewDecoder(file)
    if err := decoder.Decode(&data); err != nil {
        return
    }

    var filtered []Record
    for _, r := range data {
        if r.Active && r.Age >= 30 {
            filtered = append(filtered, r)
        }
    }

    sort.Slice(filtered, func(i, j int) bool {
        return filtered[i].Name < filtered[j].Name
    })

    output, _ := json.Marshal(filtered)
    fmt.Print(string(output))
}