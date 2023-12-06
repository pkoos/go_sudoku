package main

import (
	"testing"
	"strings"
)

func TestGridValueWithZeroValue(t * testing.T) {
	test_value := 0
	test_column := 0
	expected_string := "<td class=\"col_0\"></td>"
	test_string := gridValue(test_value, test_column)
	if strings.Compare(expected_string, test_string) != 0 {
		t.Fatalf("\nexpected: %s\nactual:   %s", expected_string, test_string)
	}
}

func TestGridValueWithNonZeroValue(t * testing.T) {
	test_value := 5
	test_column := 0
	expected_string := "<td class=\"col_0\">5</td>"
	test_string := gridValue(test_value, test_column)
	if strings.Compare(expected_string, test_string) != 0 {
		t.Fatalf("\nexpected: %s\nactual:   %s", expected_string, test_string)
	}
}