package integration_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
	"github.com/timsamart/code-concat/internal/processor"
)

type MockClipboard struct {
	content string
}

func (m *MockClipboard) WriteAll(text string) error {
	m.content = text
	return nil
}

func (m *MockClipboard) ReadAll() (string, error) {
	return m.content, nil
}

func TestDirectoryConcatenation(t *testing.T) {
	// Create a temporary directory for our test files
	tempDir, err := ioutil.TempDir("", "code-concat-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create test file structure
	createTestFiles(t, tempDir)

	// Create a mock clipboard
	mockClipboard := &MockClipboard{}

	// Run the tool
	p := processor.NewProcessor(1024, true, []string{"excluded"}, mockClipboard)
	content, err := p.ProcessDirectory(tempDir)
	if err != nil {
		t.Fatalf("ProcessDirectory failed: %v", err)
	}

	// Verify output
	expectedContent := `<file1.txt>
` + "```" + `
This is file 1
` + "```" + `

<file2.py>
` + "```python" + `
print('Hello, World!')
` + "```" + `

<subdir/file3.html>
` + "```html" + `
<html><body>Test</body></html>
` + "```" + `

<test.srt>
` + "```" + `
Speaker 1:
Hello
Speaker 2:
Hi
` + "```" + `


`

	// Print full expected content
	t.Logf("Expected content:\n%s", expectedContent)

	// Print full actual content
	t.Logf("Actual content:\n%s", content)

	// Save actual output to file
	err = ioutil.WriteFile("actual_output.txt", []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to write actual output: %v", err)
	}

	// Save expected output to file
	err = ioutil.WriteFile("expected_output.txt", []byte(expectedContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write expected output: %v", err)
	}

	if content != expectedContent {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(expectedContent, content, true)
		t.Errorf("Content mismatch. Diff:\n%s", dmp.DiffPrettyText(diffs))
		t.Errorf("Please check 'expected_output.txt' and 'actual_output.txt' for detailed comparison")
	}

	// Check if content was "copied" to our mock clipboard
	if mockClipboard.content != expectedContent {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(expectedContent, mockClipboard.content, true)
		t.Errorf("Clipboard content mismatch. Diff:\n%s", dmp.DiffPrettyText(diffs))
	}
}

func createTestFiles(t *testing.T, baseDir string) {
	files := map[string]string{
		"file1.txt":          "This is file 1",
		"file2.py":           "print('Hello, World!')",
		"subdir/file3.html":  "<html><body>Test</body></html>",
		"excluded/file4.txt": "This should be excluded",
		"test.srt":           "1\n00:00:01,000 --> 00:00:02,000\nSpeaker 1: Hello\n\n2\n00:00:03,000 --> 00:00:04,000\nSpeaker 2: Hi",
		"large_file.txt":     strings.Repeat("a", 2*1024*1024), // 2MB file
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
