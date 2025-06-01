package main

import (
	"os"
	"fmt"
	"strconv"

	"github.com/pehringer/mapper/internal/io"
	"github.com/pehringer/mapper/internal/ga"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("mapper [INPUT_MAPPINGS_FILEPATH] [ACCURACY] [OUTPUT_PROGRAM_FILEPATH]")
		return
	}
	inputPath := os.Args[1]
	accuracy, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Println("accuracy:", err)
		return
	} else if accuracy <= 0.0 || accuracy >= 1.0 {
		fmt.Println("accuracy: must be between 0.0 amd 1.0")
		return
	}
	outputPath := os.Args[3]
	mappings, err := io.ReadMappings(inputPath)
	if err != nil {
		panic(err)
	}
	program := ga.Evolve(mappings, accuracy, 48, (1024 * 2))
	err = io.WriteProgram(outputPath, program)
	if err != nil {
		panic(err)
	}
}
