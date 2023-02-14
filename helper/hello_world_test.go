package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	// Calling the function `HelloWorld` with the parameter `Adam` and assigning the result to the
	// variable `result`.
	result := HelloWorld("Adam")
	// Checking if the result is not equal to "Hello Adam"
	if result != "Hello Adam" {
		// Throwing an error.
		panic("Result is not Hello Adam")
	}
}
