package io

import (
	"fmt"
	"time"
)

var (
	start int64 = 0
)

func PrintStarting() {
	fmt.Print("0% 0ms")
	start = time.Now().UnixMilli()
}

func PrintProgress(fraction float64) {
	elapsed := time.Now().UnixMilli() - start
	percent := fraction * 100
	if percent < 100.0 {
		fmt.Printf("\r%.0f%% %dms", percent, elapsed)
	} else {
		fmt.Printf("\r100%% %dms", elapsed)
	}
}

func PrintComplete() {
	fmt.Println()
}
