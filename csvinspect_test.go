package main

import "testing"

func TestInspectCsv(t *testing.T) {
	result, err := InspectCsv("testdata/1.csv")
	if err != nil {
		t.Error(err)
		return
	}
	if result.RecordCount != 1 {
		t.Errorf("Expected 1 record, got %d", result.RecordCount)
	}
	if result.FieldCount != 3 {
		t.Errorf("Expected 3 fields, got %d", result.FieldCount)
	}
}
