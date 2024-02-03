package main

import (
	"fmt"
	"time"

	"clock/utils"
	"github.com/inancgumus/screen"
)

func main() {

	// For Go Playground, do not use this.
	screen.Clear()

	// Go Playground will not run an infinite loop.
	// So, instead, you may loop for 1000 times:
	// for i := 0; i < 1000; i++ {
	for {
		// For Go Playground, use this instead:
		// fmt.Print("\f")
		screen.MoveTopLeft()

		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()

		clock := [...]utils.Placeholder{
			utils.Digits[hour/10], utils.Digits[hour%10],
			utils.Separators,
			utils.Digits[min/10], utils.Digits[min%10],
			utils.Separators,
			utils.Digits[sec/10], utils.Digits[sec%10],
		}

		for line := range clock[0] {
			for index, digit := range clock {
				next := clock[index][line]
				if digit == utils.Separators && sec%2 == 0 {
					next = "   "
				}
				fmt.Print(next, "  ")
			}
			fmt.Println()
		}

		// pause for 1 second
		time.Sleep(time.Second)
	}

}
