package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Tests struct {
	name     string
	request  string
	expected string
}

// In Go, table tests are a way of testing a function with a variety of input values and expected output values.

func TestTableTest(t *testing.T) {
	// A slice of struct.
	tests := []Tests{
		{
			name:     "HelloWorld(Adam)",
			request:  "Adam",
			expected: "Hello Adam",
		},
		{
			name:     "HelloWorld(Joko)",
			request:  "Joko",
			expected: "Hello Joko",
		},
	}

	// A loop that iterates over the slice of struct.
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}
