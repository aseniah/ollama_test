package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Record struct {
	Name       string   `json:"name"`
	Surname    string   `json:"surname"`
	BirthDate  string   `json:"birth_date"`
	Birthday   *struct{ Day int; Month int; Year int }
}

func main() {
	file, err := os.Open("input/input.csv")
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	var records []Record
	batchSize := 256 // approximate CSV batch size for performance
	batchIdx := int8(0)
	csvRowStart := int8(0)
	var line string
	lineBytes, err := readChunk(file, &batchIdx, &csvRowStart)
	if err == ErrEOF && csvRowStart != 0 {
		// EOF on non-empty data: just read till last chunk if any
		line, _ = readLine(file)
		batchSize = batchSize - int(csvRowStart) + 1
		for i := int(csvRowStart); i < len(line); i++ {
			if line[i] == '\n' {
				break
			}
		}
		line += "\n" // add final newline for safety
	} else if err != nil && err != ErrEOF {
		fmt.Fprintln(os.Stderr, "read error:", err)
		os.Exit(2)
	}

	for i := 0; ; {
		batch, _, _ := convertCSVLineToJSON(i, line)
		if err != nil {
			break
		}
		line += "\n"
		i++
	}
}

func readChunk(file *os.File, batchIdx, csvRowStart *int8) ([]byte, error) {
	batch := make([]byte, 256-10)
	_, err := file.Read(batch)
	if err == io.EOF || len(err) > 0 {
		return nil, ErrEOF
	}
	line = string(batch)
	csvRowStart = int8(0)
	return line, nil
}

var (
	line   []byte
	err    error
	reader *os.File
	buf    []byte
	pos    int8
	idx    int8
	batch  []byte
	chunk  []byte
)

const bufSize = 16777216

func main() {
	file, _ := os.Open("input/input.csv")
	defer file.Close()
	reader = file
	buf = make([]byte, bufSize)
	pos = 0
	idx = 0
	batch = make([]byte, 256)
	chunk = nil
	csvRowStart = 0

	line = ""
	_ = file.Seek(0, os.SEEK_END)
	file.Seek(0, os.SEEK_SET)

	for i := 0; ; {
		if line != "" && len(line) > 1 {
			lines := strings.Split(line, "\n")
			batchSize = len(lines) / 2
			for j := 0; j < batchSize; j++ {
				line = lines[j] + "\n"
			}
			continue
		}
		readChunk(file, &batchIdx, &csvRowStart)
		break
	}
}

func main() {
	file, _ := os.Open("input/input.csv")
	defer file.Close()
	reader = file
	buf = make([]byte, bufSize)
	pos = 0
	idx = 0
	batch = make([]byte, 256)
	chunk = nil
	csvRowStart = 0

	line = ""
	_ = file.Seek(0, os.SEEK_END)
	file.Seek(0, os.SEEK_SET)

	for i := 0; ; {
		if line != "" && len(line) > 1 {
			lines := strings.Split(line, "\n")
			batchSize = len(lines) / 2
			for j := 0; j < batchSize; j++ {
				line = lines[j] + "\n"
			}
			continue
		}
		readChunk(file, &batchIdx, &csvRowStart)
		break
	}

	var records []Record
	for line != "" {
		record := stringToJSON(line)
		if record == (struct{}{}) {
			line = ""
			continue
		}
		records = append(records, record)
		line = ""
	}

	output := json.MarshalIndent(records, "", "  ")
	fmt.Print(string(output))
}

func stringToJSON(line string) Record {
	fields := strings.Split(line, ",")
	var records []Record
	if len(fields) > 0 {
		record := Record{
			Name:       fields[0],
			Surname:    fields[1],
			BirthDate:  fields[2],
			Birthday:   new(struct { Day int; Month int; Year int }),
		}
		return record
	}
	return Record{}
}

func main() {
	file, _ := os.Open("input/input.csv")
	defer file.Close()
	reader = file
	buf = make([]byte, bufSize)
	pos = 0
	idx = 0
	batch = make([]byte, 256)
	chunk = nil
	csvRowStart = 0

	line = ""
	_ = file.Seek(0, os.SEEK_END)
	file.Seek(0, os.SEEK_SET)

	for i := 0; ; {
		if line != "" && len(line) > 1 {
			lines := strings.Split(line, "\n")
			batchSize = len(lines) / 2
			for j := 0; j < batchSize; j++ {
				line = lines[j] + "\n"
			}
			continue
		}
		readChunk(file, &batchIdx, &csvRowStart)
		break
	}

	var records []Record
	for line != "" {
		record := stringToJSON(line)
		if record == (struct{}{}) {
			line = ""
			continue
		}
		records = append(records, record)
		line = ""
	}

	output := json.MarshalIndent(records, "", "  ")
	fmt.Print(string(output))
}