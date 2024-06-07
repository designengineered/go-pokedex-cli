package main

import (
	"testing"
)

// Test cleanInput function

func TestCleanInput(t *testing.T) {
	//Creating and def test cases slice of struc
	testCases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "Hello, World! 123",
			expected: []string{"Hello", "World", "123"},
		},
		{
			input:    "Go!Lang 2024.",
			expected: []string{"GoLang", "2024"},
		},
		{
			input:    "123@#$%",
			expected: []string{"123"},
		},
		{
			input:    "  spaces  ",
			expected: []string{"spaces"},
		},
		{
			input:    "",
			expected: []string{""},
		},
	}
	for _, tc := range testCases {
		result := cleanInput(tc.input)
		// checking if the result array converted the correct # of words
		if len(result) != len(tc.expected) {
			t.Errorf("Len Expected: %v got %v", len(tc.expected), len(result))
			//if this fails we don't need to check the content of the word
			continue
		}
		// checking content of current word
		for i := range result {
			resultWord := result[i]
			expectedWord := tc.expected[i]
			if resultWord != expectedWord {
				t.Errorf("Word Expected: %v, Word Received: %v", expectedWord, resultWord)
			}
		}
	}
}
