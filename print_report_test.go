package main

import (
	"reflect"
	"testing"
)

func TestSortPages(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		expected []PageVisit
	}{
		{
			name: "order count descending",
			input: map[string]int{
				"url1": 5,
				"url2": 1,
				"url3": 3,
				"url4": 10,
				"url5": 7,
			},
			expected: []PageVisit{
				{Url: "url4", Visits: 10},
				{Url: "url5", Visits: 7},
				{Url: "url1", Visits: 5},
				{Url: "url3", Visits: 3},
				{Url: "url2", Visits: 1},
			},
		},
		{
			name: "alphabetize",
			input: map[string]int{
				"d": 1,
				"a": 1,
				"e": 1,
				"b": 1,
				"c": 1,
			},
			expected: []PageVisit{
				{Url: "a", Visits: 1},
				{Url: "b", Visits: 1},
				{Url: "c", Visits: 1},
				{Url: "d", Visits: 1},
				{Url: "e", Visits: 1},
			},
		},
		{
			name: "order count then alphabetize",
			input: map[string]int{
				"d": 2,
				"a": 1,
				"e": 3,
				"b": 1,
				"c": 2,
			},
			expected: []PageVisit{
				{Url: "e", Visits: 3},
				{Url: "c", Visits: 2},
				{Url: "d", Visits: 2},
				{Url: "a", Visits: 1},
				{Url: "b", Visits: 1},
			},
		},
		{
			name:     "empty map",
			input:    map[string]int{},
			expected: []PageVisit{},
		},
		{
			name:     "nil map",
			input:    nil,
			expected: []PageVisit{},
		},
		{
			name: "one key",
			input: map[string]int{
				"url1": 1,
			},
			expected: []PageVisit{
				{Url: "url1", Visits: 1},
			},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := sortPages(tc.input)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
