package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	// Calling the function `HelloWorld` with the parameter `Adam` and assigning the result to the
	// variable `result`.
	result := HelloWorld("Adam")
	// Checking if the result is not equal to "Hello Adam"
	if result != "Hello Adam" {
		// Throwing an error.
		// Test failed, but we can keep going
		t.Errorf("Return is not Hello Adam")
	}

	// it will execute or continuing with the test
	fmt.Println("TestHelloWorld done!")
}

func TestHelloWorldAdam(t *testing.T) {
	// Calling the function `HelloWorld` with the parameter `Adam` and assigning the result to the
	// variable `result`.
	result := HelloWorld("Adam")
	// Checking if the result is not equal to "Hello Adam"
	if result != "Hello Adam" {
		// Throwing an error.
		// Test failed, and we have to stop
		t.Fatal("Return is not Hello Adam")
	}

	// This code will never be executed
	fmt.Println("TestHelloWorldAdam done!")
}
