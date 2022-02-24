package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type InspectResult struct {
	RecordCount int `json:"record_count"`
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: csvinspect <filename>")
		os.Exit(1)
	}
	csvFile := os.Args[1]
	inCh := ReadCsv(csvFile)
	result := InspectResult{}
	for readResult := range inCh {
		if readResult.Error != nil {
			fmt.Println("Error:", readResult.Error)
			os.Exit(1)
		}
		result.RecordCount++
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
