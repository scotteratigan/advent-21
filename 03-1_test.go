package main

import (
	"reflect"
	"strconv"
	"testing"
)

func TestBinStrToInt(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"0", 0},
		{"1", 1},
		{"10", 2},
		{"11", 3},
		{"111", 7},
		{"1001", 9},
	}

	for _, tc := range tests {
		result := BinStrToInt(tc.input)
		if result != tc.expected {
			t.Errorf("Wrong value, expected " + strconv.Itoa(tc.expected) + " but received " + strconv.Itoa(result))
		}
	}

}

func TestGetDigitCounts(t *testing.T) {
	inputs := []string{"111", "001", "000"}
	expected := []count{
		{
			zeros: 2,
			ones:  1,
		},
		{
			zeros: 2,
			ones:  1,
		},
		{
			zeros: 1,
			ones:  2,
		},
	}
	result := GetDigitCounts(inputs)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Incorrect digit count result.")
	}
}
