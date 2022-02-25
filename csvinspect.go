package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type InspectResult struct {
	// The number of records
	RecordCount int `json:"record_count"`
	// The number of fields of the first record
	FieldCount int `json:"field_count"`
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: csvinspect <filename>")
		os.Exit(1)
	}
	csvFile := os.Args[1]
	result, err := InspectCsv(csvFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	jsonResult, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println(string(jsonResult))
}

type ReadResult struct {
	Record []string
	Error  error
}

func InspectCsv(csvFile string) (*InspectResult, error) {
	result := InspectResult{}
	inCh := ReadCsv(csvFile)
	for readResult := range inCh {
		if readResult.Error != nil {
			return nil, readResult.Error
		}
		if result.RecordCount == 0 {
			result.FieldCount = len(readResult.Record)
		}
		result.RecordCount++
	}
	return &result, nil
}

func ReadCsv(csvFile string) chan ReadResult {
	ch := make(chan ReadResult)
	go func() {
		defer close(ch)
		// open csvFile
		f, err := os.Open(csvFile)
		if err != nil {
			ch <- ReadResult{nil, err}
			return
		}
		defer f.Close()
		// create a new csv reader
		r := csv.NewReader(f)
		for {
			record, err := r.Read()
			if err != nil {
				if err == io.EOF {
					break
				}
				ch <- ReadResult{nil, err}
				return
			}
			ch <- ReadResult{record, nil}
		}
	}()
	return ch
}
