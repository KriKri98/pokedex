package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{{
		input:    "  hello  world ",
		expected: []string{"hello", "world"},
	},
		{
			input:    "this is a test",
			expected: []string{"this", "is", "a", "test"},
		},
		{
			input:    "I LIKE  TRAINS",
			expected: []string{"i", "like", "trains"},
		},
		{
			input:    "You Are Weird",
			expected: []string{"you", "are", "weird"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Fail: expected length %v, actual length %v", len(c.expected), len(actual))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Fail: expected %s, got %s", expectedWord, word)
			}
		}
	}

}
