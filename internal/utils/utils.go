package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func ValidateDirectory(path string) error {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return fmt.Errorf("directory %s does not exist", path)
	}
	if !info.IsDir() {
		return fmt.Errorf("path %s is not a directory", path)
	}
	return nil
}

func IsGitIgnored(path string) bool {
	cmd := exec.Command("git", "check-ignore", "-q", path)
	cmd.Dir = filepath.Dir(path)
	err := cmd.Run()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			if exitErr.ExitCode() == 128 {
				return false
			}
		}
		return false
	}
	return true
}

func IsText(bytes []byte) bool {
	const sampleSize = 512
	if len(bytes) < sampleSize {
		return isPrintableChars(bytes)
	}
	return isPrintableChars(bytes[:sampleSize])
}

func isPrintableChars(data []byte) bool {
	for _, b := range data {
		if b != '\n' && b != '\r' && b != '\t' && (b < 32 || b > 126) {
			return false
		}
	}
	return true
}
