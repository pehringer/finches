package io

import (
	"fmt"
	"time"
)

var (
	start int64 = 0
)

func PrintStarting() {
	fmt.Print("running...")
	start = time.Now().UnixMilli()
}

func PrintProgress(fraction float64) {
	elapsed := time.Now().UnixMilli() - start
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
		fmt.Printf("%.0f%% %dms", percent, elapsed)
	} else {
		fmt.Printf("100%% %dms", elapsed)
	}
}

func PrintComplete() {
	fmt.Println()
}
