package template_test

import (
	"fmt"

	"github.com/mrz1836/go-template"
)

// ExampleGreet demonstrates the usage of the Greet function.
func ExampleGreet() {
	msg := template.Greet("Alice")
	fmt.Println(msg)
	// Output: Hello Alice
}
