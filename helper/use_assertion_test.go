package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// The function UseAssertion takes two integers as input and returns an integer
func TestUseAssertion(t *testing.T) {
	// we store value from UseAssertion to result
	result := UseAssertion(2, 2)
	// compare between result and value that we expect
	// because we use assert as compare function, testify will invoke Fail() function
	assert.Equal(t, 43, result)
	// it will executed
	fmt.Println("Dieksekusi")
}

func TestUseRequire(t *testing.T) {
	// we store value from UseAssertion to result
	result := UseAssertion(2, 2)
	// compare between result and value that we expect
	// because we use require as compare function, testify will invoke FailNow() function
	require.Equal(t, 42, result)
	// it will not executed
	fmt.Println("Dieksekusi")
}
