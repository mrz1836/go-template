//go:build mage

// Magefile for go-broadcast specific tasks
package main

import (
	"fmt"

	"github.com/magefile/mage/sh"
)

// BenchHeavy runs intensive benchmarks excluded from CI
// This may take 10-30 minutes and includes:
// - Worker pool stress tests (1000+ concurrent tasks)
// - Large directory sync simulations
// - Memory efficiency tests with large datasets
// - Real-world scenario simulations
func BenchHeavy() error {
	return sh.RunV("go", "test", "-bench=.", "-benchmem",
		"-tags=bench_heavy", "-benchtime=1s", "-timeout=60m", "./...")
}

// BenchAll runs all benchmarks including heavy ones
// This may take 30-60 minutes total
func BenchAll() error {
	// Run default benchmarks first
	if err := sh.RunV("go", "test", "-bench=.", "-benchmem",
		"-benchtime=100ms", "-timeout=20m", "./..."); err != nil {
		return fmt.Errorf("quick benchmarks failed: %w", err)
	}

	// Then run heavy benchmarks
	return BenchHeavy()
}

// BenchQuick runs only the quick benchmarks (same as magex bench)
func BenchQuick() error {
	return sh.RunV("go", "test", "-bench=.", "-benchmem",
		"-benchtime=100ms", "-timeout=20m", "./...")
}

// TestQuick runs fast unit tests excluding performance tests
func TestQuick() error {
	return sh.RunV("go", "test", "-short", "./...")
}

// TestPerf runs performance regression tests with build tag
// These tests are excluded from regular runs due to long execution time
func TestPerf() error {
	return sh.RunV("go", "test", "-tags=performance", "-timeout=30m", "./test/integration")
}

// TestAll runs all tests including performance tests
func TestAll() error {
	if err := TestQuick(); err != nil {
		return fmt.Errorf("quick tests failed: %w", err)
	}
	return TestPerf()
}
