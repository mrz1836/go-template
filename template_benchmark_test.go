package template_test

import (
	"testing"

	"github.com/mrz1836/go-template"
)

// BenchmarkGreet benchmarks the Greet function.
func BenchmarkGreet(b *testing.B) {
	for b.Loop() {
		_ = template.Greet("BenchmarkUser")
	}
}
