package bigfilehandler

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// FileType represents the type of the file content to return
type FileType int

const (
	// TEXT returns the file content as plain text
	TEXT FileType = iota
	// JSON returns the file content as a JSON object
	JSON
)

// ReadFile reads a big file and returns its content as either text or JSON
func ReadFile(filePath string, fileType FileType) (interface{}, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	switch fileType {
	case TEXT:
		content, err := readAsText(file)
		if err != nil {
			return nil, err
		}
		return content, nil

	case JSON:
		content, err := readAsJSON(file)
		if err != nil {
			return nil, err
		}
		return content, nil

	default:
		return nil, fmt.Errorf("unsupported file type")
	}
}

func readAsText(file *os.File) (string, error) {
	scanner := bufio.NewScanner(file)
	var content string
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}
	return content, nil
}

func readAsJSON(file *os.File) (interface{}, error) {
	decoder := json.NewDecoder(file)
	var content interface{}
	if err := decoder.Decode(&content); err != nil {
		return nil, fmt.Errorf("error decoding JSON: %w", err)
	}
	return content, nil
}
