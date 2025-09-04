//go:build mage

package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidatePath(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{
			name:    "valid relative path",
			path:    "test/file.txt",
			wantErr: false,
		},
		{
			name:    "valid current directory file",
			path:    "./file.txt",
			wantErr: false,
		},
		{
			name:    "invalid path traversal with ../",
			path:    "../outside/file.txt",
			wantErr: true,
		},
		{
			name:    "invalid path traversal with ../../",
			path:    "../../outside/file.txt",
			wantErr: true,
		},
		{
			name:    "cleaned path that becomes invalid",
			path:    "safe/../../../etc/passwd",
			wantErr: true,
		},
		{
			name:    "valid nested path",
			path:    "deep/nested/path/file.go",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validatePath(tt.path)
			if tt.wantErr {
				require.Error(t, err)
				assert.Equal(t, errInvalidPath, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestParseArgument(t *testing.T) {
	tests := []struct {
		name            string
		arg             string
		expectedOwner   string
		expectedRepo    string
		expectedDryRun  bool
		expectedVerbose bool
		expectedCleanup bool
	}{
		{
			name:          "parse owner argument",
			arg:           "owner=testowner",
			expectedOwner: "testowner",
		},
		{
			name:         "parse repo argument",
			arg:          "repo=testrepo",
			expectedRepo: "testrepo",
		},
		{
			name:           "parse dryrun true",
			arg:            "dryrun=true",
			expectedDryRun: true,
		},
		{
			name:            "parse verbose true",
			arg:             "verbose=true",
			expectedVerbose: true,
		},
		{
			name:            "parse cleanup true",
			arg:             "cleanup=true",
			expectedCleanup: true,
		},
		{
			name: "ignore unknown argument",
			arg:  "unknown=value",
		},
		{
			name:          "parse owner with special characters",
			arg:           "owner=test-owner_123",
			expectedOwner: "test-owner_123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var owner, repo string

			var dryRun, verbose, cleanup bool

			parseArgument(tt.arg, &owner, &repo, &dryRun, &verbose, &cleanup)

			assert.Equal(t, tt.expectedOwner, owner)
			assert.Equal(t, tt.expectedRepo, repo)
			assert.Equal(t, tt.expectedDryRun, dryRun)
			assert.Equal(t, tt.expectedVerbose, verbose)
			assert.Equal(t, tt.expectedCleanup, cleanup)
		})
	}
}

func TestValidateRequiredArgs(t *testing.T) {
	tests := []struct {
		name    string
		owner   string
		repo    string
		wantErr error
	}{
		{
			name:  "valid args",
			owner: "testowner",
			repo:  "testrepo",
		},
		{
			name:    "missing owner",
			owner:   "",
			repo:    "testrepo",
			wantErr: errMissingOwner,
		},
		{
			name:    "missing repo",
			owner:   "testowner",
			repo:    "",
			wantErr: errMissingRepo,
		},
		{
			name:    "missing both",
			owner:   "",
			repo:    "",
			wantErr: errMissingOwner, // Should return first missing arg error
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateRequiredArgs(tt.owner, tt.repo)
			if tt.wantErr != nil {
				assert.Equal(t, tt.wantErr, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestIsBinaryFile(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected bool
	}{
		{
			name:     "text file",
			path:     "main.go",
			expected: false,
		},
		{
			name:     "png image",
			path:     "image.png",
			expected: true,
		},
		{
			name:     "jpeg image",
			path:     "photo.jpg",
			expected: true,
		},
		{
			name:     "executable",
			path:     "program.exe",
			expected: true,
		},
		{
			name:     "shared library",
			path:     "lib.so",
			expected: true,
		},
		{
			name:     "case insensitive",
			path:     "IMAGE.PNG",
			expected: true,
		},
		{
			name:     "markdown file",
			path:     "README.md",
			expected: false,
		},
		{
			name:     "no extension",
			path:     "Makefile",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isBinaryFile(tt.path)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestApplyReplacements(t *testing.T) {
	tests := []struct {
		name         string
		content      string
		replacements []struct{ from, to string }
		path         string
		verbose      bool
		wantContent  string
		wantModified bool
	}{
		{
			name:    "single replacement",
			content: "package mrz1836/go-template",
			replacements: []struct{ from, to string }{
				{"mrz1836/go-template", "testowner/testrepo"},
			},
			path:         "test.go",
			wantContent:  "package testowner/testrepo",
			wantModified: true,
		},
		{
			name:    "multiple replacements",
			content: "mrz1836/go-template and go-template by mrz1836",
			replacements: []struct{ from, to string }{
				{"mrz1836/go-template", "testowner/testrepo"},
				{"go-template", "testrepo"},
				{"mrz1836", "testowner"},
			},
			path:         "test.go",
			wantContent:  "testowner/testrepo and testrepo by testowner",
			wantModified: true,
		},
		{
			name:    "no replacements needed",
			content: "some random content",
			replacements: []struct{ from, to string }{
				{"notfound", "replacement"},
			},
			path:         "test.go",
			wantContent:  "some random content",
			wantModified: false,
		},
		{
			name:         "empty content",
			content:      "",
			replacements: []struct{ from, to string }{{"test", "replace"}},
			path:         "empty.go",
			wantContent:  "",
			wantModified: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotContent, gotModified := applyReplacements(
				tt.content, tt.replacements, tt.path, tt.verbose,
			)
			assert.Equal(t, tt.wantContent, gotContent)
			assert.Equal(t, tt.wantModified, gotModified)
		})
	}
}

func TestCreateReplacements(t *testing.T) {
	owner := "testowner"
	repo := "testrepo"

	replacements := createReplacements(owner, repo)

	expected := []struct{ from, to string }{
		{"mrz1836/go-template", "testowner/testrepo"},
		{"go-template", "testrepo"},
		{"mrz1836", "testowner"},
	}

	assert.Equal(t, expected, replacements)
}

func TestShouldSkipFile(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		isDir    bool
		expected bool
	}{
		{
			name:     "regular go file",
			path:     "main.go",
			isDir:    false,
			expected: false,
		},
		{
			name:     "git directory",
			path:     ".git/config",
			isDir:    true,
			expected: true,
		},
		{
			name:     "idea directory",
			path:     ".idea/workspace.xml",
			isDir:    true,
			expected: true,
		},
		{
			name:     "make directory",
			path:     ".make/build",
			isDir:    true,
			expected: true,
		},
		{
			name:     "binary file",
			path:     "image.png",
			isDir:    false,
			expected: true,
		},
		{
			name:     "regular directory",
			path:     "src/",
			isDir:    true,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock FileInfo
			var info os.FileInfo
			if tt.isDir {
				info = &mockFileInfo{name: filepath.Base(tt.path), isDir: true}
			} else {
				info = &mockFileInfo{name: filepath.Base(tt.path), isDir: false}
			}

			result := shouldSkipFile(tt.path, info)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestProcessFile_PathValidation(t *testing.T) {
	// Test that processFile properly validates paths
	replacements := []struct{ from, to string }{
		{"test", "replacement"},
	}

	// Test invalid path
	modified, err := processFile("../invalid/path", replacements, true, false)
	require.Error(t, err)
	assert.False(t, modified)
	assert.Contains(t, err.Error(), "path validation failed")
}

// Mock implementation of os.FileInfo for testing
type mockFileInfo struct {
	name  string
	isDir bool
}

func (m *mockFileInfo) Name() string       { return m.name }
func (m *mockFileInfo) Size() int64        { return 0 }
func (m *mockFileInfo) Mode() os.FileMode  { return 0o644 }
func (m *mockFileInfo) ModTime() time.Time { return time.Time{} }
func (m *mockFileInfo) IsDir() bool        { return m.isDir }
func (m *mockFileInfo) Sys() interface{}   { return nil }

// Test the parseInstallArgs function with mocked os.Args
func TestParseInstallArgsWithMockedArgs(t *testing.T) {
	tests := []struct {
		name          string
		args          []string
		expectedOwner string
		expectedRepo  string
		expectedDry   bool
		expectedVerb  bool
		expectedClean bool
		expectError   bool
		errorType     error
	}{
		{
			name:          "valid arguments",
			args:          []string{"cmd", "owner=testowner", "repo=testrepo"},
			expectedOwner: "testowner",
			expectedRepo:  "testrepo",
		},
		{
			name:        "missing owner",
			args:        []string{"cmd", "repo=testrepo"},
			expectError: true,
			errorType:   errMissingOwner,
		},
		{
			name:        "missing repo",
			args:        []string{"cmd", "owner=testowner"},
			expectError: true,
			errorType:   errMissingRepo,
		},
		{
			name:          "with flags",
			args:          []string{"cmd", "owner=testowner", "repo=testrepo", "dryrun=true", "verbose=true", "cleanup=true"},
			expectedOwner: "testowner",
			expectedRepo:  "testrepo",
			expectedDry:   true,
			expectedVerb:  true,
			expectedClean: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Save original args and restore after test
			originalArgs := os.Args

			defer func() { os.Args = originalArgs }()

			// Set test args
			os.Args = tt.args

			owner, repo, dryRun, verbose, cleanup, err := parseInstallArgs()

			if tt.expectError {
				require.Error(t, err)
				assert.Equal(t, tt.errorType, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedOwner, owner)
				assert.Equal(t, tt.expectedRepo, repo)
				assert.Equal(t, tt.expectedDry, dryRun)
				assert.Equal(t, tt.expectedVerb, verbose)
				assert.Equal(t, tt.expectedClean, cleanup)
			}
		})
	}
}

// Benchmark tests for performance-critical functions
func BenchmarkValidatePath(b *testing.B) {
	paths := []string{
		"valid/path/file.go",
		"../invalid/path",
		"./current/dir/file.txt",
		"deeply/nested/path/structure/file.go",
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, path := range paths {
			_ = validatePath(path)
		}
	}
}

func BenchmarkIsBinaryFile(b *testing.B) {
	files := []string{
		"main.go",
		"image.png",
		"document.pdf",
		"library.so",
		"README.md",
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, file := range files {
			_ = isBinaryFile(file)
		}
	}
}

func BenchmarkApplyReplacements(b *testing.B) {
	content := strings.Repeat("mrz1836/go-template is a template by mrz1836 for go-template projects. ", 100)
	replacements := []struct{ from, to string }{
		{"mrz1836/go-template", "testowner/testrepo"},
		{"go-template", "testrepo"},
		{"mrz1836", "testowner"},
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = applyReplacements(content, replacements, "benchmark.go", false)
	}
}
