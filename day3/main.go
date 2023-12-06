package day3

import (
	"fmt"
	"math"
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

func bound(x, min, max int) int {
	return int(math.Min(math.Max(float64(min), float64(x)), float64(max)))
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
				if y == maxYBound || !isNum(string(line[y])) {
					foundNum, _ := strconv.Atoi(curNum)
					total += foundNum
					curNum = ""
					allowedNums[x] = append(allowedNums[x], foundNum)
					foundAdj = false
				}
			}

			if curNum == "" {
				continue
			}

			if curNum == "4" {
				utils.LogDebug("found 454")
			}

			if isNum(curNum) && !isNum(c) {
				curNum = ""
			}

			for xCursor := (x - 1); xCursor < (x + 2); xCursor += 1 {
				for yCursor := (y - 1); yCursor < (y + 2); yCursor += 1 {
					xPos := bound(xCursor, minXBound, maxXBound)
					yPos := bound(yCursor, minYBound, maxYBound)
					adjChar := string(lines[xPos][yPos])
					if !re.Match([]byte(adjChar)) {
						foundAdj = true
					}
				}
			}
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
