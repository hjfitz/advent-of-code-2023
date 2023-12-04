package utils

import (
	"fmt"
	"os"
)

func LogDebug(format string, args ...interface{}) {
	debug := os.Getenv("DEBUG")
	if debug == "true" || debug == "1" {
		fmt.Printf(format, args...)
	}
}
