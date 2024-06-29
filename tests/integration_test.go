package integration_test

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/timsamart/code-concat/cmd/dircopier"
)

func TestDirectoryConcatenation(t *testing.T) {
	// Create a temporary directory for our test files
	tempDir, err := ioutil.TempDir("", "code-concat-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create test file structure
	createTestFiles(t, tempDir)

	// Run the tool
	var output bytes.Buffer
	processor := dircopier.NewProcessor(1024, true, []string{"excluded"})
	err = processor.ProcessDirectory(tempDir, &output)
	if err != nil {
		t.Fatalf("ProcessDirectory failed: %v", err)
	}

	// Verify output
	result := output.String()

	// Check if all non-excluded files are included
	expectedFiles := []string{"file1.txt", "file2.py", "subdir/file3.html"}
	for _, file := range expectedFiles {
		if !strings.Contains(result, file) {
			t.Errorf("Output doesn't contain expected file: %s", file)
		}
	}

	// Check if excluded directory is not included
	if strings.Contains(result, "excluded/file4.txt") {
		t.Errorf("Output contains file from excluded directory")
	}

	// Check if SRT file is processed correctly
	if !strings.Contains(result, "Speaker 1:") || strings.Contains(result, "00:00:01,000") {
		t.Errorf("SRT file not processed correctly")
	}

	// Check if file size limit is respected
	if strings.Contains(result, "large_file.txt") {
		t.Errorf("Output contains large file that should have been skipped")
	}
}

func createTestFiles(t *testing.T, baseDir string) {
	files := map[string]string{
		"file1.txt":          "This is file 1",
		"file2.py":           "print('Hello, World!')",
		"subdir/file3.html":  "<html><body>Test</body></html>",
		"excluded/file4.txt": "This should be excluded",
		"test.srt":           "1\n00:00:01,000 --> 00:00:02,000\nSpeaker 1: Hello\n\n2\n00:00:03,000 --> 00:00:04,000\nSpeaker 2: Hi",
		"large_file.txt":     strings.Repeat("a", 2048),
	}

	for path, content := range files {
		fullPath := filepath.Join(baseDir, path)
		err := os.MkdirAll(filepath.Dir(fullPath), 0755)
		if err != nil {
			t.Fatalf("Failed to create directory: %v", err)
		}
		err = ioutil.WriteFile(fullPath, []byte(content), 0644)
		if err != nil {
			t.Fatalf("Failed to write file: %v", err)
		}
	}
}
