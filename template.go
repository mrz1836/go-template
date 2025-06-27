// Package template provides a robust starter template for building new Go libraries.
//
// This package implements foundational patterns and utilities for Go library development and is designed to help developers quickly scaffold, test, and maintain secure, idiomatic Go code.
//
// Key features include:
// - Built-in support for code quality, testing, and CI/CD workflows
// - Example functions and best-practice patterns for Go libraries
//
// The package is structured for modularity and ease of extension, following Go community conventions. It relies on the Go standard library and popular tools for testing and linting.
//
// Usage examples:
//
//	msg := template.Greet("Alice")
//	fmt.Println(msg) // Output: Hello Alice
//
// Important notes:
// - Assumes Go modules are used for dependency management
// - No external configuration is required for basic usage
// - Designed for use as a starting point for new Go projects
//
// This package is part of the go-template project and is intended to be copied or forked for new Go library development.
package template

import "fmt"

// Greet returns a greeting message for the given first name.
//
// This function performs the following steps:
// - Formats a greeting string using the provided first name.
//
// Parameters:
// - firstname: The first name to include in the greeting message.
//
// Returns:
// - A string containing the greeting message.
//
// Side Effects:
// - None.
//
// Notes:
// - Assumes firstname is a non-empty string; no validation is performed.
// - This function is standalone and not part of a larger workflow.
func Greet(firstname string) string {
	return fmt.Sprintf("Hello %s", firstname)
}
