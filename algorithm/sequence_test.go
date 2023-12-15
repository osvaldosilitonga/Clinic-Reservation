package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindSequence(t *testing.T) {
	m := []int{20, 7, 8, 10, 2, 5, 6}

	type testCase struct {
		main     []int
		seq      []int
		expected bool
	}

	// Test Case
	cases := []testCase{
		{
			main:     m,
			seq:      []int{7, 8},
			expected: true,
		},
		{
			main:     m,
			seq:      []int{8, 7},
			expected: false,
		},
		{
			main:     m,
			seq:      []int{7, 10},
			expected: false,
		},
	}

	for _, c := range cases {
		result := FindSequence(c.main, c.seq)
		assert.Equal(t, c.expected, result, "FindSequence was incorrect, got: %t, want: %t.", result, c.expected)
	}
}
