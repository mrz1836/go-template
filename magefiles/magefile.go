//go:build mage

// Magefile for go-template specific tasks
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/magefile/mage/sh"
)

var (
	errMissingOwner           = errors.New("missing required parameter: owner=<value>")
	errMissingRepo            = errors.New("missing required parameter: repo=<value>")
	errInvalidPath            = errors.New("invalid path: path traversal not allowed")
	errUnexpectedModuleFormat = errors.New("unexpected module path format")
	errModuleNotFound         = errors.New("module declaration not found in go.mod")
)

// logMessage prints a message to stdout (allowed alternative to fmt.Printf)
func logMessage(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stdout, format, args...)
}

// logLine prints a line to stdout (allowed alternative to fmt.Println)
func logLine(msg string) {
	_, _ = fmt.Fprintf(os.Stdout, "%s\n", msg)
}

// validatePath checks if a file path is safe to read
func validatePath(path string) error {
	// Clean the path to prevent path traversal
	cleanPath := filepath.Clean(path)

	// Ensure the path doesn't escape current directory
	if strings.Contains(cleanPath, "..") {
		return errInvalidPath
	}

	return nil
}

// parseArgument processes a single command line argument
func parseArgument(arg string, owner, repo *string, dryRun, verbose, cleanup *bool) {
	switch {
	case strings.HasPrefix(arg, "owner="):
		*owner = strings.TrimPrefix(arg, "owner=")
	case strings.HasPrefix(arg, "repo="):
		*repo = strings.TrimPrefix(arg, "repo=")
	case arg == "dryrun=true":
		*dryRun = true
	case arg == "verbose=true":
		*verbose = true
	case arg == "cleanup=true":
		*cleanup = true
	}
}

// validateRequiredArgs checks that required parameters are provided
func validateRequiredArgs(owner, repo string) error {
	if owner == "" {
		return errMissingOwner
	}

	if repo == "" {
		return errMissingRepo
	}

	return nil
}

// parseInstallArgs extracts parameters from command line arguments or MAGE_ARGS environment variable
func parseInstallArgs() (owner, repo string, dryRun, verbose, cleanup bool, err error) {
	var args []string

	// Check if MAGE_ARGS is available (mage binary execution)
	if mageArgs := os.Getenv("MAGE_ARGS"); mageArgs != "" {
		// Use MAGE_ARGS environment variable (for mage binary execution)
		args = strings.Fields(mageArgs)
	} else if len(os.Args) > 2 {
		// Use os.Args (for go run execution), skip the command name
		args = os.Args[2:] // Skip program name and function name
	}

	for _, arg := range args {
		parseArgument(arg, &owner, &repo, &dryRun, &verbose, &cleanup)
	}

	if err = validateRequiredArgs(owner, repo); err != nil {
		return "", "", false, false, false, err
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

// applyReplacements applies string replacements to content
func applyReplacements(content string, replacements []struct{ from, to string }, path string, verbose bool) (string, bool) {
	newContent := content
	modified := false

	for _, replacement := range replacements {
		if strings.Contains(newContent, replacement.from) {
			newContent = strings.ReplaceAll(newContent, replacement.from, replacement.to)
			modified = true

			if verbose {
				logMessage("   📝 %s: %s → %s\n", path, replacement.from, replacement.to)
			}
		}
	}

	return newContent, modified
}

// writeFileIfModified writes content to file if changes were made and not in dry run
func writeFileIfModified(path, content string, modified, dryRun bool) error {
	if modified && !dryRun {
		if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
			return err
		}
	}

	return nil
}

// processFile performs string replacements in a text file
func processFile(path string, replacements []struct{ from, to string }, dryRun, verbose bool) (bool, error) {
	// Validate path to prevent path traversal attacks
	if err := validatePath(path); err != nil {
		return false, fmt.Errorf("path validation failed for %s: %w", path, err)
	}

	content, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return false, err
	}

	originalContent := string(content)
	newContent, modified := applyReplacements(originalContent, replacements, path, verbose)

	if err = writeFileIfModified(path, newContent, modified, dryRun); err != nil {
		return false, err
	}

	return modified, nil
}

// getTemplateInfo extracts the current module owner and repo name from go.mod
func getTemplateInfo() (owner, repo string, err error) {
	file, err := os.Open("go.mod")
	if err != nil {
		return "", "", fmt.Errorf("failed to open go.mod: %w", err)
	}

	defer func() {
		_ = file.Close()
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "module ") {
			modulePath := strings.TrimPrefix(line, "module ")
			// Expected format: github.com/owner/repo or similar
			parts := strings.Split(modulePath, "/")
			if len(parts) >= 3 {
				owner = parts[len(parts)-2]
				repo = parts[len(parts)-1]

				return owner, repo, nil
			}

			return "", "", fmt.Errorf("%w: %s", errUnexpectedModuleFormat, modulePath)
		}
	}

	if err := scanner.Err(); err != nil {
		return "", "", fmt.Errorf("error reading go.mod: %w", err)
	}

	return "", "", errModuleNotFound
}

// getFallbackTemplateInfo returns template info by trying multiple locations
// This is extracted to make it testable and avoid hardcoding values in multiple places
func getFallbackTemplateInfo() (owner, repo string, err error) {
	// Try to extract from go.mod in current directory
	owner, repo, err = getTemplateInfo()
	if err == nil {
		return owner, repo, nil
	}

	// If not found, try parent directory (magefiles subdirectory case)
	originalDir, dirErr := os.Getwd()
	if dirErr != nil {
		return "", "", fmt.Errorf("failed to get working directory: %w", dirErr)
	}

	if chErr := os.Chdir(".."); chErr != nil {
		return "", "", fmt.Errorf("failed to change to parent directory: %w", chErr)
	}

	defer func() {
		_ = os.Chdir(originalDir)
	}()

	owner, repo, err = getTemplateInfo()
	if err != nil {
		return "", "", fmt.Errorf("failed to read template info from go.mod: %w", err)
	}

	return owner, repo, nil
}

// createReplacements creates the list of string replacements needed
func createReplacements(owner, repo string) []struct{ from, to string } {
	// Get current template owner and repo from go.mod
	templateOwner, templateRepo, err := getTemplateInfo()
	if err != nil {
		// Fallback: try multiple locations (current dir and parent dir)
		// This allows the function to work when called from subdirectories
		templateOwner, templateRepo, err = getFallbackTemplateInfo()
		if err != nil {
			// If we still can't read go.mod, log error and continue with empty values
			// This shouldn't happen in normal usage, but makes the code more resilient
			logMessage("Warning: Could not read template info from go.mod: %v\n", err)
		}
	}

	return []struct {
		from string
		to   string
	}{
		{fmt.Sprintf("%s/%s", templateOwner, templateRepo), fmt.Sprintf("%s/%s", owner, repo)},
		{templateRepo, repo},
		{templateOwner, owner},
	}
}

// shouldSkipFile determines if a file should be skipped during processing
func shouldSkipFile(path string, info os.FileInfo) bool {
	if info.IsDir() {
		return strings.HasPrefix(path, ".git/") || strings.HasPrefix(path, ".idea/") || strings.HasPrefix(path, ".make/")
	}

	return isBinaryFile(path)
}

// processFileInWalk processes a single file during walk
func processFileInWalk(path string, info os.FileInfo, replacements []struct{ from, to string }, dryRun, verbose bool, modifiedFiles *[]string) error {
	if info.IsDir() {
		if shouldSkipFile(path, info) {
			return filepath.SkipDir
		}

		return nil
	}

	if shouldSkipFile(path, info) {
		return nil
	}

	modified, err := processFile(path, replacements, dryRun, verbose)
	if err != nil {
		return fmt.Errorf("error processing %s: %w", path, err)
	}

	if modified {
		*modifiedFiles = append(*modifiedFiles, path)
	}

	return nil
}

// processAllFiles walks through all files and applies replacements
func processAllFiles(replacements []struct{ from, to string }, dryRun, verbose bool) ([]string, error) {
	var modifiedFiles []string

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		return processFileInWalk(path, info, replacements, dryRun, verbose, &modifiedFiles)
	})

	return modifiedFiles, err
}

// removeShareImage removes the default share image if it exists
func removeShareImage(shareImagePath string, verbose bool) ([]string, error) {
	var cleanupFiles []string

	if _, err := os.Stat(shareImagePath); err != nil {
		// File doesn't exist, nothing to remove
		return cleanupFiles, nil //nolint:nilerr // File not existing is acceptable
	}

	if err := os.Remove(shareImagePath); err != nil {
		return nil, fmt.Errorf("error removing share image: %w", err)
	}

	cleanupFiles = append(cleanupFiles, shareImagePath+" (removed)")

	if verbose {
		logMessage("🗑️  Removed default share image\n")
	}

	return cleanupFiles, nil
}

// handleSelfCleanup displays messages about self-cleanup
func handleSelfCleanup(cleanup, verbose bool) {
	if !cleanup {
		return
	}

	if verbose {
		logLine("🧹 Self-cleanup requested - removing InstallTemplate function")
	}

	logLine("⚠️  Manual cleanup required: Remove InstallTemplate function from magefile.go")
}

// performCleanupOperations handles cleanup tasks after file processing
func performCleanupOperations(dryRun, verbose, cleanup bool) ([]string, error) {
	if dryRun {
		return []string{}, nil
	}

	shareImagePath := ".github/IMAGES/go-share-image.png"

	cleanupFiles, err := removeShareImage(shareImagePath, verbose)
	if err != nil {
		return nil, err
	}

	handleSelfCleanup(cleanup, verbose)

	return cleanupFiles, nil
}

// reportResults displays the final results to the user
func reportResults(modifiedFiles []string, dryRun bool) {
	logMessage("✅ Template initialization complete!\n")
	logMessage("📁 Modified %d files:\n", len(modifiedFiles))

	for _, file := range modifiedFiles {
		logMessage("   • %s\n", file)
	}

	if !dryRun {
		logMessage("\n🎯 Next steps:\n")
		logMessage("   1. Review all changes: git diff\n")
		logMessage("   2. Test your new project: go test ./...\n")
		logMessage("   3. Commit initial version: git add . && git commit -m \"Initial commit from go-template\"\n")
	}
}

// InstallTemplate initializes a new project from go-template
// Usage: magex InstallTemplate owner=yourorg repo=yourrepo [dryrun=true] [verbose=true] [cleanup=true]
func InstallTemplate() error {
	owner, repo, dryRun, verbose, cleanup, err := parseInstallArgs()
	if err != nil {
		return fmt.Errorf("parameter error: %w", err)
	}

	if verbose {
		logMessage("🚀 Initializing template for %s/%s...\n", owner, repo)

		if dryRun {
			logLine("🔍 Dry run mode - no changes will be made")
		}
	}

	replacements := createReplacements(owner, repo)

	modifiedFiles, err := processAllFiles(replacements, dryRun, verbose)
	if err != nil {
		return fmt.Errorf("error walking files: %w", err)
	}

	cleanupFiles, err := performCleanupOperations(dryRun, verbose, cleanup)
	if err != nil {
		return err
	}

	allModifiedFiles := append(modifiedFiles, cleanupFiles...)
	reportResults(allModifiedFiles, dryRun)

	return nil
}

// TestQuick runs fast unit tests excluding performance tests
func TestQuick() error {
	return sh.RunV("go", "test", "-short", "./...")
}
