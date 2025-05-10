package io

import (
	"fmt"
)

var (
	spinIndex int = 0
	spinChars string = "|/-\\"
)

func PrintProgress(fraction float64) {
	percent := fraction * 100
	bars := int(percent / 2)
	fmt.Print("\r")
	for i := 0; i < 50; i++ {
		if i < bars {
			fmt.Print("|")
		} else {
			fmt.Print(" ")
		}
	}
	if percent < 100.0 {
		fmt.Printf("%.0f%%", percent)
		fmt.Printf(" %c ", spinChars[spinIndex%len(spinChars)])
		spinIndex++
	} else {
		fmt.Print("100%  ")
	}
}

func PrintComplete() {
	fmt.Println("\n------------------------mapper------------------------")
}

