package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/hjfitz/advent-of-code-2023/day1"
)

func main() {
	days := []func(){
		day1.Run,
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
		for _, fn := range days {
			fn()
		}
	}
}
