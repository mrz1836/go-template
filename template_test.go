package template_test

import (
	"testing"

	"github.com/mrz1836/go-template"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestGreet tests the Greet function with various input scenarios using table-driven tests.
func TestGreet(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "normal name",
			input:    "Alice",
			expected: "Hello Alice",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "Hello ",
		},
		{
			name:     "whitespace",
			input:    " ",
			expected: "Hello  ",
		},
		{
			name:     "special characters",
			input:    "@!$",
			expected: "Hello @!$",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := template.Greet(tt.input)
			require.NotNil(t, result)
			assert.Equal(t, tt.expected, result)
		})
	}
}
