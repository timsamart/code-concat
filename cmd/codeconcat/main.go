package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/timsamart/code-concat/internal/processor"
	"github.com/timsamart/code-concat/internal/utils"
)

var (
	directoryPath string
	maxSizeKB     int
	processSRT    bool
	excludedDirs  []string
)

func init() {
	flag.BoolVar(&processSRT, "srt", false, "Process SRT files to clean timestamps and group by speaker")
	flag.IntVar(&maxSizeKB, "size", 1024, "Maximum file size in KB to process")
	flag.IntVar(&maxSizeKB, "s", 1024, "Maximum file size in KB (shorthand)")
	flag.Func("exclude", "Directories to exclude", func(s string) error {
		excludedDirs = strings.Split(s, ",")
		return nil
	})
	flag.Func("e", "Directories to exclude (shorthand)", func(s string) error {
		excludedDirs = strings.Split(s, ",")
		return nil
	})

	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Println("Usage: codeconcat [flags] <directoryPath>")
		os.Exit(1)
	}

	directoryPath = flag.Arg(0)
}

func main() {
	if err := utils.ValidateDirectory(directoryPath); err != nil {
		log.Fatalf("Directory validation failed: %v", err)
	}

	clipboardWriter := &clipboardWrapper{}
	p := processor.NewProcessor(maxSizeKB, processSRT, excludedDirs, clipboardWriter)
	content, err := p.ProcessDirectory(directoryPath)
	if err != nil {
		log.Fatalf("Error processing directory: %v", err)
	}

	if err := clipboard.WriteAll(content); err != nil {
		log.Fatalf("Failed to copy content to clipboard: %v", err)
	}

	fmt.Println("Content copied to clipboard.")
}

// clipboardWrapper implements the ClipboardWriter interface
type clipboardWrapper struct{}

func (cw *clipboardWrapper) WriteAll(text string) error {
	return clipboard.WriteAll(text)
}
