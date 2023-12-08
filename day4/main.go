package day4

import (
	"math"
	"regexp"
	"strconv"
	"strings"

	//"strings"

	"github.com/hjfitz/advent-of-code-2023/utils"
)

const SAMPLE_DATA = `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

// tuple return, [winners, numbers]
func parseCard(card string) [][]int {
	result := [][]int{}
	re := regexp.MustCompile("Card [0-9]+: ")
	cardRaw := re.ReplaceAll([]byte(card), []byte(""))
	parts := strings.Split(string(cardRaw), " | ")

	for _, part := range parts {
		partArr := []int{}
		nums := strings.Split(part, " ")
		for _, num := range nums {
			curNum := strings.TrimSpace(num)
			if curNum != "" {
				parsedVal, _ := strconv.Atoi(strings.TrimSpace(num))
				partArr = append(partArr, parsedVal)
			}
		}
		result = append(result, partArr)
	}

	utils.LogDebug("Raw card: %s\n", cardRaw)

	return result
}

func calcMatches(game, winningNumbers []int) int {
	utils.LogDebug("%+v\n", game)
	utils.LogDebug("%+v\n", winningNumbers)
	total := 0
	// go greedy
	for _, ourPoint := range game {
		for _, winner := range winningNumbers {
			if ourPoint == winner {
				utils.LogDebug("%d ", winner)
				total += 1
			}
		}
	}
	utils.LogDebug("\n")

	return total
}

func calcPoints(won int) int {
	if won <= 0 {
		return 0
	}
	return int(math.Pow(2, float64(won-1)))
}

func sum(arr []int) (total int) {
	for _, val := range arr {
		total += val
	}
	return
}

func Part1(lines []string) int {
	results := []int{}
	for _, line := range lines {
		utils.LogDebug("%+v\n", line)
		parsedCard := parseCard(line)
		winners := parsedCard[0]
		game := parsedCard[1]
		matches := calcMatches(game, winners)
		points := calcPoints(matches)

		results = append(results, points)

	}
	return sum(results)
}

func Part2(lines []string) int {
	results := []int{}
	res := map[int]int{}
	for i, line := range lines {
		utils.LogDebug("%+v\n", line)
		parsedCard := parseCard(line)
		winners := parsedCard[0]
		game := parsedCard[1]
		matches := calcMatches(game, winners)
		end := i + matches
		for i := i; i < end; i += 1 {
			res[i] += 1
		}
		utils.LogDebug("index: %d, matches: %d, end: %d\n", i, matches, end)
		//results = append(results, Part2(lines[i:i+matches]))

	}
	for k, _ := range lines {
		v := res[k]

		utils.LogDebug("k: %d, v: %d\n", k+1, v)
	}
	return sum(results)
}

func Run() {
	parts := []func([]string) int{
		Part1,
		Part2, // broken
	}
	utils.Runner(4, SAMPLE_DATA, parts, true)
}
