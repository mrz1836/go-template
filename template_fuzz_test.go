package template_test

import (
	"testing"

	"github.com/mrz1836/go-template"
	"github.com/stretchr/testify/require"
)

// FuzzGreet validates that Greet always returns a string that starts with "Hello " and is at least 6 characters long.
func FuzzGreet(f *testing.F) {
	seed := []string{
		"Alice",
		"",
		"123",
		"!@#",
		"世界", //nolint: gosmopolitan // Test with non-ASCII characters
	}
	for _, tc := range seed {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, input string) {
		out := template.Greet(input)
		require.GreaterOrEqualf(t, len(out), 6, "output too short: %q", out)
		require.Equalf(t, "Hello ", out[:6], "output does not start with 'Hello ': %q", out)
	})
}
