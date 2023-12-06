package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/hjfitz/advent-of-code-2023/day1"
	"github.com/hjfitz/advent-of-code-2023/day2"
	"github.com/hjfitz/advent-of-code-2023/day3"
)

func main() {
	days := []func(){
		day1.Run,
		day2.Run,
		day3.Run,
	}

	if len(os.Args) > 1 {
		dayNumRaw := os.Args[1]
		dayNum, _ := strconv.Atoi(dayNumRaw)

		idx := dayNum - 1
		fmt.Printf("Running day %d\n", dayNum)

		fn := days[idx]
		fn()
	} else {
		fmt.Println("Running all days")
		fmt.Println()
		for idx, fn := range days {
			day := idx + 1
			fmt.Printf("Running day %d\n", day)
			fn()
			fmt.Println()
		}
	}
}
