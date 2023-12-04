package day1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const INPUT_DATA_PART_ONE = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

const INPUT_DATA_PART_TWO = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

func getLines(filename string) []string {
	data, _ := os.ReadFile(filename)
	lines := strings.Split(string(data), "\n")
	return lines
}

func getValues(line string) int {
	pair := ""
	vals := []string{}
	for _, r := range line {
		char := string(r)
		_, err := strconv.Atoi(char)
		// no error, we were able to convert successfully
		if err == nil {
			vals = append(vals, char)
		}
	}
	if len(vals) > 0 {
		pair = pair + vals[0]
		pair = pair + vals[len(vals)-1]
	} else {
		pair = "0"
	}
	val, _ := strconv.Atoi(pair)
	return val
}

func partTwoPerLine(line string) int {
	lookup := map[string]rune{
		//"zero":  '0',
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}

	allowedValues := []string{
		"one", "1",
		"two", "2",
		"three", "3",
		"four", "4",
		"five", "5",
		"six", "6",
		"seven", "7",
		"eight", "8",
		"nine", "9",
	}
	start := ""
	end := ""

	// some dirty linear searching begins
searchPrefix:
	for i := 0; i < len(line); i += 1 {
		window := line[i:]
		for _, a := range allowedValues {
			if strings.HasPrefix(window, a) {
				start = a
				break searchPrefix
			}
		}
	}

	// then from the back
searchSuffix:
	for i := 0; i < len(line); i += 1 {
		window := line[0 : len(line)-i]
		for _, a := range allowedValues {
			if strings.HasSuffix(window, a) {
				end = a
				break searchSuffix
			}
		}
	}

	for k, v := range lookup {
		if start == k {
			start = string(v)
		}
		if end == k {
			end = string(v)
		}
	}


	val := start + end

	parsed, _ := strconv.Atoi(val)
	return parsed
}

func Day1Part1(lines []string) int {
	total := 0
	for _, line := range lines {
		pair := getValues(line)
		total += pair
	}
	return total
}

func Day1Part2(lines []string) int {
	total := 0
	for _, line := range lines {
		val := partTwoPerLine(line)
		total += val
	}
	return total
}

func Run() {
	sampleDataPartTwo := strings.Split(INPUT_DATA_PART_TWO, "\n")
	actualData := getLines("day1/day1")
	sampleData := strings.Split(INPUT_DATA_PART_ONE, "\n")

	partOneSample := Day1Part1(sampleData)
	fmt.Printf("Day 1, Part 1 for sample data: %d\n", partOneSample)

	partOneAnswer := Day1Part1(actualData)
	fmt.Printf("Day 1, Part 1 for actual data: %d\n", partOneAnswer)

	partTwoSample := Day1Part2(sampleDataPartTwo)
	fmt.Printf("Day 1, Part 2 for sample data: %d\n", partTwoSample)

	partTwoAnswer := Day1Part2(actualData)
	fmt.Printf("Day 1, Part 2 for actual data: %d\n", partTwoAnswer)
}
