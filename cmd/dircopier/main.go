package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

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
		fmt.Println("Usage: dircopier [flags] <directoryPath>")
		os.Exit(1)
	}

	directoryPath = flag.Arg(0)
}

func main() {
	if err := utils.ValidateDirectory(directoryPath); err != nil {
		log.Fatalf("Directory validation failed: %v", err)
	}

	p := processor.NewProcessor(maxSizeKB, processSRT, excludedDirs)
	if err := p.ProcessDirectory(directoryPath); err != nil {
		log.Fatalf("Error processing directory: %v", err)
	}

	fmt.Println("Content copied to clipboard.")
}
