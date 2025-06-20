package main

import (
	"os"
	"fmt"
	"strconv"
	"github.com/pehringer/mapper/internal/io"
	"github.com/pehringer/mapper/internal/ga"
)

func main() {
	if len(os.Args) % 2 != 0 {
		fmt.Println("mapper /path/to/mappings.csv")
		fmt.Println("Options:")
		fmt.Println("\t-gen [GENERATION_COUNT]")
		fmt.Println("\t-pop [POPULATION_COUNT]")
		return
	}
	filepath := os.Args[1]
	generations := 1024
	population := 1024
	for i := 2; i < len(os.Args); i += 2 {
		number, err := strconv.ParseInt(os.Args[i+1], 10, 64)
		if err != nil {
			panic(err)
		}
		if os.Args[i] == "-gen" {
			generations = int(number)
		}
		if os.Args[i] == "-pop" {
			population = int(number)
		}
	}
	inputs, outputs, err := io.ReadMappings(filepath)
	if err != nil {
		panic(err)
	}
	constants, instructions := ga.Evolution(inputs, outputs, generations, population)
	err = io.WriteSolution("program.asm", constants, instructions)
	if err != nil {
		panic(err)
	}
}
