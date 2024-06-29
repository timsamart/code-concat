package filehandler

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"
)

func HandleFile(buffer *bytes.Buffer, relPath string, content []byte, processSRT bool) {
	fileExt := filepath.Ext(relPath)
	buffer.WriteString(fmt.Sprintf("<%s>\n", relPath))
	buffer.WriteString("```")
	buffer.WriteString(getLanguageSpecifier(fileExt))
	buffer.WriteString("\n")

	if fileExt == ".srt" && processSRT {
		buffer.WriteString(strings.TrimSpace(cleanSRT(string(content))))
	} else {
		buffer.WriteString(strings.TrimSpace(string(content)))
	}

	buffer.WriteString("\n```\n\n")
}

func getLanguageSpecifier(ext string) string {
	switch ext {
	case ".py":
		return "python"
	case ".js":
		return "javascript"
	case ".html":
		return "html"
	case ".css":
		return "css"
	case ".go":
		return "go"
	// Add more cases as needed
	default:
		return "" // No language specifier for unknown extensions
	}
}

func cleanSRT(content string) string {
	var result strings.Builder
	lines := strings.Split(content, "\n")
	currentSpeaker := ""
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.Contains(line, "-->") || (len(line) > 0 && line[0] >= '0' && line[0] <= '9') {
			continue
		}
		if strings.Contains(line, ":") {
			parts := strings.SplitN(line, ":", 2)
			currentSpeaker = parts[0]
			result.WriteString(currentSpeaker + ":\n")
			line = strings.TrimSpace(parts[1])
		}
		result.WriteString(line + "\n")
	}
	return strings.TrimSpace(result.String())
}
