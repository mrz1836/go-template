//go:build mage

// Magefile for go-template specific tasks
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/magefile/mage/sh"
)

// parseInstallArgs extracts parameters from command line arguments
func parseInstallArgs() (owner, repo string, dryRun, verbose, cleanup bool, err error) {
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "owner=") {
			owner = strings.TrimPrefix(arg, "owner=")
		} else if strings.HasPrefix(arg, "repo=") {
			repo = strings.TrimPrefix(arg, "repo=")
		} else if arg == "dryrun=true" {
			dryRun = true
		} else if arg == "verbose=true" {
			verbose = true
		} else if arg == "cleanup=true" {
			cleanup = true
		}
	}

	if owner == "" {
		return "", "", false, false, false, fmt.Errorf("missing required parameter: owner=<value>")
	}
	if repo == "" {
		return "", "", false, false, false, fmt.Errorf("missing required parameter: repo=<value>")
	}

	return owner, repo, dryRun, verbose, cleanup, nil
}

// isBinaryFile checks if a file is binary based on its extension
func isBinaryFile(path string) bool {
	binaryExts := []string{
		".png", ".jpg", ".jpeg", ".gif", ".ico", ".pdf", ".zip", ".tar", ".gz",
		".exe", ".dll", ".so", ".dylib", ".bin", ".dat",
	}

	ext := strings.ToLower(filepath.Ext(path))
	for _, binaryExt := range binaryExts {
		if ext == binaryExt {
			return true
		}
	}
	return false
}

// processFile performs string replacements in a text file
func processFile(path string, replacements []struct{ from, to string }, dryRun, verbose bool) (bool, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return false, err
	}

	originalContent := string(content)
	newContent := originalContent
	modified := false

	// Apply all replacements
	for _, replacement := range replacements {
		if strings.Contains(newContent, replacement.from) {
			newContent = strings.ReplaceAll(newContent, replacement.from, replacement.to)
			modified = true
			if verbose {
				fmt.Printf("   📝 %s: %s → %s\n", path, replacement.from, replacement.to)
			}
		}
	}

	// Write back if modified and not dry run
	if modified && !dryRun {
		if err = os.WriteFile(path, []byte(newContent), 0o644); err != nil {
			return false, err
		}
	}

	return modified, nil
}

// InstallTemplate initializes a new project from go-template
// Usage: magex installTemplate owner=yourorg repo=yourrepo [dryrun=true] [verbose=true] [cleanup=true]
func InstallTemplate() error {
	owner, repo, dryRun, verbose, cleanup, err := parseInstallArgs()
	if err != nil {
		return fmt.Errorf("parameter error: %w", err)
	}

	if verbose {
		fmt.Printf("🚀 Initializing template for %s/%s...\n", owner, repo)
		if dryRun {
			fmt.Println("🔍 Dry run mode - no changes will be made")
		}
	}

	var modifiedFiles []string

	// Perform replacements
	replacements := []struct {
		from string
		to   string
	}{
		{"mrz1836/go-template", fmt.Sprintf("%s/%s", owner, repo)},
		{"go-template", repo},
		{"mrz1836", owner},
	}

	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories and specific paths
		if info.IsDir() {
			if strings.HasPrefix(path, ".git/") || strings.HasPrefix(path, ".idea/") || strings.HasPrefix(path, ".make/") {
				return filepath.SkipDir
			}
			return nil
		}

		// Skip binary files
		if isBinaryFile(path) {
			return nil
		}

		// Process text files
		var modified bool
		modified, err = processFile(path, replacements, dryRun, verbose)
		if err != nil {
			return fmt.Errorf("error processing %s: %w", path, err)
		}

		if modified {
			modifiedFiles = append(modifiedFiles, path)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("error walking files: %w", err)
	}

	// Cleanup operations
	if !dryRun {
		shareImagePath := ".github/IMAGES/go-share-image.png"
		if _, err = os.Stat(shareImagePath); err == nil {
			if err = os.Remove(shareImagePath); err != nil {
				return fmt.Errorf("error removing share image: %w", err)
			}
			modifiedFiles = append(modifiedFiles, shareImagePath+" (removed)")
			if verbose {
				fmt.Printf("🗑️  Removed default share image\n")
			}
		}

		// Optional self-cleanup
		if cleanup {
			if verbose {
				fmt.Println("🧹 Self-cleanup requested - removing InstallTemplate function")
			}
			// Note: This would require more complex implementation to modify this file
			fmt.Println("⚠️  Manual cleanup required: Remove InstallTemplate function from magefile.go")
		}
	}

	// Report results
	fmt.Printf("✅ Template initialization complete!\n")
	fmt.Printf("📁 Modified %d files:\n", len(modifiedFiles))
	for _, file := range modifiedFiles {
		fmt.Printf("   • %s\n", file)
	}

	if !dryRun {
		fmt.Printf("\n🎯 Next steps:\n")
		fmt.Printf("   1. Review all changes: git diff\n")
		fmt.Printf("   2. Test your new project: go test ./...\n")
		fmt.Printf("   3. Commit initial version: git add . && git commit -m \"Initial commit from go-template\"\n")
	}

	return nil
}

// TestQuick runs fast unit tests excluding performance tests
func TestQuick() error {
	return sh.RunV("go", "test", "-short", "./...")
}
