package utils

import (
	"fmt"
	"strings"
)

func Runner(day int, sampleData string, parts []func([]string) int, sampleOnly bool) {
	sampleLines := strings.Split(strings.TrimSpace(sampleData), "\n")
	lines := GetLines(fmt.Sprintf("files/day%d", day))
	for idx, part := range parts {
		partNum := idx + 1
		partSample := part(sampleLines)
		fmt.Printf("Day %d, Part %d for sample data: %d\n", day, partNum, partSample)
		if !sampleOnly {
			partActual := part(lines)
			fmt.Printf("Day %d, Part %d for actual data: %d\n", day, partNum, partActual)
		}
	}

}
