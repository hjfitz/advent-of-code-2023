package utils

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func FileExists(filePath string) bool {
	info, err := os.Stat(filePath)
	if err == nil {
		return !info.IsDir()
	}
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return false
}

func GetLines(filename string) []string {
	if !FileExists(filename) {
		fmt.Printf("[WARN] File '%s' does not exist\n", filename)
		return []string{}
	}
	data, _ := os.ReadFile(filename)
	trimmed := strings.TrimSpace(string(data))
	lines := strings.Split(trimmed, "\n")
	return lines
}
