package utils

import (
	"os"
	"strings"
)

func GetLines(filename string) []string {
	data, _ := os.ReadFile(filename)
	trimmed := strings.TrimSpace(string(data))
	lines := strings.Split(trimmed, "\n")
	return lines
}
