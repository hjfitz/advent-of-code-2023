package day3

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/hjfitz/advent-of-code-2023/utils"
)

const SAMPLE_DEBUG_DATA = `
....820....*.................#...................707...
...../.................118&.....922.....785........./..
.........../......753..................................`

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

func isNum(c string) bool {
	re := regexp.MustCompile("[0-9]")
	return re.Match([]byte(c))
}

func Part1(lines []string) (total int) {
	re := regexp.MustCompile("[0-9]|\\.")
	allowedNums := [][]int{}
	for x := 0; x < len(lines); x += 1 {
		allowedNums = append(allowedNums, []int{})
		line := lines[x]
		curNum := ""
		foundAdj := false
		for y := 0; y < len(line); y += 1 {

			c := string(line[y])
			utils.LogDebug("%s", c)

			minXBound := 0
			maxXBound := len(lines) - 1
			minYBound := 0
			maxYBound := len(line) - 1

			if isNum(c) {
				curNum += c
			}

			if foundAdj && curNum != "" {
				// when the number has ended, it either
				// has a non-numeric to the right: !isNum(line[y+1])
				// or is at the end (y == maxYBound)
				// if y == maxYBound || (y < maxYBound && !isNum(string(line[y+1]))) {
				if y == maxYBound || !isNum(string(line[y])) {
					foundNum, _ := strconv.Atoi(curNum)
					total += foundNum
					curNum = ""
					allowedNums[x] = append(allowedNums[x], foundNum)
					foundAdj = false
				}
			}

			// if we don't have a number, skip checking adjacent characters
			if curNum == "" {
				continue
			}

			// if we do have a number, but the number has ended, remove that number
			if isNum(curNum) && !isNum(c) {
				curNum = ""
			}

			// check above
			if x > minXBound {
				adjChar := string(lines[x-1][y])

				if !re.Match([]byte(adjChar)) {
					foundAdj = true
				}

				// utils.LogDebug("'u: %s' ", adjChar)
			}

			// check left
			if y > minYBound {
				adjChar := string(line[y-1])

				if !re.Match([]byte(adjChar)) {
					foundAdj = true
				}

				//utils.LogDebug("'l: %s' ", adjChar)
			}

			// check below
			if x < maxXBound {
				adjChar := string(lines[x+1][y])

				if !re.Match([]byte(adjChar)) {
					foundAdj = true
				}
				//utils.LogDebug("'b: %s' ", adjChar)
			}

			// check right
			if y < maxYBound {
				adjChar := string(line[y+1])
				if !re.Match([]byte(adjChar)) {
					foundAdj = true
				}
				//utils.LogDebug("'r: %s' ", adjChar)
			}

			// check up and left (diagonal)
			if x > minXBound && y > minYBound {
				adjChar := string(lines[x-1][y-1])
				if !re.Match([]byte(adjChar)) {
					foundAdj = true
				}
				//utils.LogDebug("'ul: %s' ", adjChar)
			}

			// check up and right (other diagonal)
			if x > minXBound && y < maxYBound {
				adjChar := string(lines[x-1][y+1])

				if !re.Match([]byte(adjChar)) {
					foundAdj = true
				}

				//utils.LogDebug("'ur: %s' ", adjChar)
			}
			// check down and right (other other diagonal)
			if x < maxXBound && y < maxYBound {
				adjChar := string(lines[x+1][y+1])
				if !re.Match([]byte(adjChar)) {
					foundAdj = true
				}

				//utils.LogDebug("'bl: %s'", string(lines[x+1][y+1]))
			}

			// check down and left
			if x < maxXBound && y > minYBound {
				adjChar := string(lines[x+1][y-1])
				if !re.Match([]byte(adjChar)) {
					foundAdj = true
				}
				//utils.LogDebug("'br: %s'", string(adjChar))
			}

			foundAdj = foundAdj && curNum != ""

		}
		utils.LogDebug("\n")
	}
	for _, arr := range allowedNums {
		utils.LogDebug("%d\n", arr)
	}

	return
}

func Run() {
	sampleLines := strings.Split(strings.TrimSpace(SAMPLE_DATA), "\n")
	part1Sample := Part1(sampleLines)

	actualLines := utils.GetLines("files/day3")
	part1Actual := Part1(actualLines)

	fmt.Printf("Day 3, Part 1 for sample data: %d\n", part1Sample)
	fmt.Printf("Day 3, Part 1 for actual data: %d\n", part1Actual)
}
