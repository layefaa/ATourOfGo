// 1. PACKAGE AND IMPORTS
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

// 2. DEFINE CUSTOM ERRORS
// In Go, we define errors as variables to reuse them.
var ErrEmptyInput = errors.New("input cannot be empty")
var ErrTooShort = errors.New("input is too short (min 8 chars)")

// 3. THE VALIDATION FUNCTION
// Note the signature: (string, error).
// In Go, we return the data AND the error status simultaneously.
func validateInput(raw string) (string, error) {
	// STEP A: Clean the input using strings.TrimSpace

	// STEP B: If string is empty, return "", ErrEmptyInput

	// STEP C: If len(cleaned) < 8, return "", ErrTooShort

	// STEP D: Return cleaned, nil (nil means no error occurred)
	return "", nil
}

// 4. THE POINTER FUNCTION (State Modification)
// This function doesn't return anything. It "reaches back" to the
// original memory address to change the value.
func sanitizeInPlace(input *string) {
	// SECURITY CHECK: Always check if a pointer is nil before dereferencing!
	if input == nil {
		return
	}

	// STEP: Convert the value at the address to lowercase
	// Hint: *input = strings.ToLower(*input)
}

// 5. MAIN EXECUTION
func main() {
	// STEP A: Set up a "Reader" for the terminal
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter data to sanitize: ")

	// STEP B: Read until the newline character
	rawInput, _ := reader.ReadString('\n')

	// STEP C: THE GO "IF" PATTERN
	// This is the most important part of Go security.
	// Capture the return AND check the error immediately.
	clean, err := validateInput(rawInput)
	if err != nil {
		fmt.Printf("Validation failed: %v\n", err)
		os.Exit(1) // Fail closed
	}

	// STEP D: USE THE POINTER
	// Pass the address of 'clean' using the & operator
	sanitizeInPlace(&clean)

	// STEP E: FINAL OUTPUT
	fmt.Printf("Safe Output: %s\n", clean)
}
