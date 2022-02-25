package main

import "testing"

func TestInspectCsv1(t *testing.T) {
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
	if result.Charset != "ASCII" {
		t.Errorf("Expected ASCII, got %s", result.Charset)
	}
}

func TestInspectCsv2(t *testing.T) {
	result, err := InspectCsv("testdata/2.csv")
	if err != nil {
		t.Error(err)
		return
	}
	if result.RecordCount != 1 {
		t.Errorf("Expected 1 record, got %d", result.RecordCount)
	}
	if result.FieldCount != 2 {
		t.Errorf("Expected 2 fields, got %d", result.FieldCount)
	}
	if result.Charset != "UTF-8" {
		t.Errorf("Expected UTF-8, got %s", result.Charset)
	}
}
