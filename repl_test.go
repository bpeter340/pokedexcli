package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  helloworld  ",
			expected: []string{"helloworld"},
		},
		{
			input:    "  Charmander Bulbasaur PIKACHU ",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, c := range cases {

		actual := cleanInput(c.input)

		expectedLength, actualLength := len(c.expected), len(actual)

		if actualLength != expectedLength {
			t.Errorf("slices are not the same length: want %v got %v", expectedLength, actualLength)
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("slices are not equivalent: want %v got %v", c.expected, actual)
			}

		}
	}
}
