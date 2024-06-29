package filehandler

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"github.com/yourusername/directory-copier/internal/utils"
)

func HandleFile(buffer *bytes.Buffer, path string, content []byte, processSRT bool) {
	if utils.IsText(content) {
		if processSRT && strings.HasSuffix(path, ".srt") {
			cleanedContent := cleanSRT(string(content))
			buffer.WriteString(fmt.Sprintf("<%s>\n```\n%s\n```\n\n", path, cleanedContent))
		} else if strings.HasSuffix(path, ".py") {
			buffer.WriteString(fmt.Sprintf("<%s>\n```python\n%s\n```\n\n", path, content))
		} else if strings.HasSuffix(path, ".html") {
			buffer.WriteString(fmt.Sprintf("<%s>\n```html\n%s\n```\n\n", path, content))
		} else {
			buffer.WriteString(fmt.Sprintf("<%s>\n```\n%s\n```\n\n", path, content))
		}
		log.Printf("Processed file: %s\n", path)
	} else {
		log.Printf("Skipping non-text file: %s\n", path)
	}
}

func cleanSRT(content string) string {
	var result strings.Builder
	lines := strings.Split(content, "\n")
	currentSpeaker := ""
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.Contains(line, "-->") || line[0] >= '0' && line[0] <= '9' {
			continue
		}
		if strings.Contains(line, ":") {
			parts := strings.SplitN(line, ":", 2)
			currentSpeaker = parts[0]
			result.WriteString(currentSpeaker + ":\n")
			line = parts[1]
		}
		result.WriteString(line + "\n")
	}
	return result.String()
}
