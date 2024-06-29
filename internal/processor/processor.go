package processor

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/atotto/clipboard"
	"github.com/yourusername/directory-copier/internal/filehandler"
	"github.com/yourusername/directory-copier/internal/utils"
)

type Processor struct {
	maxSizeKB    int
	processSRT   bool
	excludedDirs []string
}

func NewProcessor(maxSizeKB int, processSRT bool, excludedDirs []string) *Processor {
	return &Processor{
		maxSizeKB:    maxSizeKB,
		processSRT:   processSRT,
		excludedDirs: excludedDirs,
	}
}

func (p *Processor) ProcessDirectory(directoryPath string) error {
	maxSizeBytes := p.maxSizeKB * 1024
	var allContent bytes.Buffer

	err := filepath.Walk(directoryPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if p.shouldSkip(path, info, int64(maxSizeBytes)) {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		content, err := ioutil.ReadFile(path)
		if err != nil {
			log.Printf("Error reading file %s: %s\n", path, err)
			return nil
		}
		filehandler.HandleFile(&allContent, path, content, p.processSRT)
		return nil
	})

	if err != nil {
		return fmt.Errorf("error walking the path: %v", err)
	}

	if err := clipboard.WriteAll(allContent.String()); err != nil {
		return fmt.Errorf("failed to copy content to clipboard: %v", err)
	}

	log.Println("Content copied to clipboard successfully.")
	return nil
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
		if utils.IsGitIgnored(path) {
			log.Printf("Skipping Git ignored directory: %s\n", path)
			return true
		}
		return false
	}
	if info.Size() > maxSizeBytes {
		log.Printf("Skipping file due to size limit: %s\n", path)
		return true
	}
	if utils.IsGitIgnored(path) {
		log.Printf("Skipping Git ignored file: %s\n", path)
		return true
	}
	return false
}
