package processor

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/timsamart/code-concat/internal/filehandler"
	"github.com/timsamart/code-concat/internal/utils"
)

type ClipboardWriter interface {
	WriteAll(string) error
}

type Processor struct {
	maxSizeKB       int
	processSRT      bool
	excludedDirs    []string
	clipboardWriter ClipboardWriter
}

func NewProcessor(maxSizeKB int, processSRT bool, excludedDirs []string, clipboardWriter ClipboardWriter) *Processor {
	return &Processor{
		maxSizeKB:       maxSizeKB,
		processSRT:      processSRT,
		excludedDirs:    excludedDirs,
		clipboardWriter: clipboardWriter,
	}
}

func (p *Processor) ProcessDirectory(directoryPath string) (string, error) {
	maxSizeBytes := p.maxSizeKB * 1024
	var allContent bytes.Buffer

	err := filepath.Walk(directoryPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			// Log the error but continue processing
			log.Printf("Error accessing path %q: %v\n", path, err)
			return nil
		}
		if p.shouldSkip(path, info, int64(maxSizeBytes)) {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		if !info.IsDir() {
			content, err := ioutil.ReadFile(path)
			if err != nil {
				log.Printf("Error reading file %s: %s\n", path, err)
				return nil
			}

			// Get relative path
			relPath, err := filepath.Rel(directoryPath, path)
			if err != nil {
				return err
			}
			// Replace backslashes with forward slashes for consistency
			relPath = filepath.ToSlash(relPath)

			filehandler.HandleFile(&allContent, relPath, content, p.processSRT)
		}
		return nil
	})

	if err != nil {
		return "", fmt.Errorf("error walking the path: %v", err)
	}

	processedContent := allContent.String()

	// Normalize line endings to LF
	processedContent = strings.ReplaceAll(processedContent, "\r\n", "\n")

	// Add an extra newline at the end
	processedContent += "\n"

	if p.clipboardWriter != nil {
		if err := p.clipboardWriter.WriteAll(processedContent); err != nil {
			return "", fmt.Errorf("failed to copy content to clipboard: %v", err)
		}
		log.Println("Content copied to clipboard successfully.")
	}

	return processedContent, nil
}

func (p *Processor) shouldSkip(path string, info os.FileInfo, maxSizeBytes int64) bool {
	if info.IsDir() {
		dirName := filepath.Base(path)
		for _, exDir := range p.excludedDirs {
			if dirName == exDir {
				log.Printf("Skipping excluded directory: %s\n", path)
				return true
			}
		}
		if dirName == ".git" {
			log.Printf("Skipping .git directory: %s\n", path)
			return true
		}
		isIgnored := utils.IsGitIgnored(path)
		log.Printf("Git ignored check for directory %s: %v", path, isIgnored)
		if isIgnored {
			log.Printf("Skipping Git ignored directory: %s\n", path)
			return true
		}
		return false
	}
	if info.Size() > maxSizeBytes {
		log.Printf("Skipping file due to size limit: %s\n", path)
		return true
	}
	isIgnored := utils.IsGitIgnored(path)
	log.Printf("Git ignored check for file %s: %v", path, isIgnored)
	if isIgnored {
		log.Printf("Skipping Git ignored file: %s\n", path)
		return true
	}
	return false
}
