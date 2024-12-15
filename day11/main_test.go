package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	input    string
	expected int
}

func TestPart1(t *testing.T) {
	cases := []TestCase{
		{input: "125 17", expected: 55312},
	}

	for _, c := range cases {
		result := part1(c.input)
		assert.Equal(t, c.expected, result, "Not expected result")
	}
}
