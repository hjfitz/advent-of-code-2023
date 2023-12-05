package day3

import (
	"fmt"
	"strconv"
	"strings"
)

const SAMPLE_DATA = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

func isNum(c byte) bool {
	_, err := strconv.Atoi(string(c))
	if err != nil {
		return false
	}
	return true
}

func Part1(lines []string) (total int) {
	for x := 0; x < len(lines); x += 1 {
		line := lines[x]
		curNum := ""
		foundAdj := false
		for y := 0; y < len(line); y += 1 {
			c := line[y]
			if isNum(c) {
				curNum += string(c)
			} else {
				if foundAdj {
					foundNum, _ := strconv.Atoi(curNum)
					total += foundNum
					curNum = ""
					foundAdj = false
				}
			}

			// check above
			if x > 0 {
			}

			// check left
			if y > 0 {
			}

			// check up and left (diagonal)
			if x > 0 && y > 0 {
			}

			// check below
			if x < (len(lines) - 1) {
			}

			// check right
			if y < (len(line) - 1) {
			}

		}
	}

	return
}

func Run() {
	sampleLines := strings.Split(SAMPLE_DATA, "\n")

	part1Sample := Part1(sampleLines)

	fmt.Printf("Day 3, Part 1 for sample data: %d\n", part1Sample)
}
