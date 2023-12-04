package day2

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/hjfitz/advent-of-code-2023/utils"
)

const SAMPLE_DATA = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

func parseLine(game string) [][]int {
	re := regexp.MustCompile("^Game [0-9]*: ")
	// remove `Game N:`
	game = string(re.ReplaceAll([]byte(game), []byte("")))
	revealsRaw := strings.Split(game, "; ")
	revs := [][]int{}
	for _, r := range revealsRaw {
		utils.LogDebug("reveal: %s\n", r)
		h := strings.Split(r, ", ")
		rev := make([]int, 3)
		for _, g := range h {
			spl := strings.Split(g, " ")
			n, _ := strconv.Atoi(string(spl[0]))
			c := spl[1]
			utils.LogDebug("num: \"%d\", col: \"%s\"\n", n, c)
			if c == "red" {
				rev[0] = n
			} else if c == "green" {
				rev[1] = n
			} else if c == "blue" {
				rev[2] = n
			}
		}
		revs = append(revs, rev)

	}

	return revs

}

func isValidGame(reveals [][]int) bool {
	const MAX_RED = 12
	const MAX_BLUE = 14
	const MAX_GREEN = 13
	for _, reveal := range reveals {
		r := reveal[0]
		g := reveal[1]
		b := reveal[2]
		if r > MAX_RED || g > MAX_GREEN || b > MAX_BLUE {
			return false
		}
	}
	return true
}

func Day2Part1(lines []string) int {
	total := 0
	for idx, l := range lines {
		utils.LogDebug("Parsing: %s\n", l)
		p := parseLine(l)
		v := isValidGame(p)
		if v {
			total += (idx + 1)
		}
		utils.LogDebug("Line: %+v, valid: %v\n", p, v)
	}
	return total
}

func Run() {
	sampleLines := strings.Split(SAMPLE_DATA, "\n")
	part1Sample := Day2Part1(sampleLines)
	part1Lines := utils.GetLines("files/day2")
	part1Actual := Day2Part1(part1Lines)

	fmt.Printf("Day 2, Part 1 for sample data: %d\n", part1Sample)
	fmt.Printf("Day 2, Part 1 for actual data: %d\n", part1Actual)
}
