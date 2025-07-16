package main

import (
	"fmt"
	"os"
	"strconv"
)

var (
	optionG int    = 10000
	optionI string = "input.csv"
	optionO string = "output.py"
	optionP int    = 1000
)

func parseGenerations(argument string) {
	value, err := strconv.ParseInt(argument, 10, 32)
	if err != nil {
		fmt.Println("[-g / --generations] must be a integer number")
		os.Exit(-1)
	}
	if value <= 0 {
		fmt.Println("[-g / --generations] must be greater than 0")
		os.Exit(-1)
	}
	optionG = int(value)
}

func parsePopulation(argument string) {
	value, err := strconv.ParseInt(argument, 10, 32)
	if err != nil {
		fmt.Println("[-p / --population] must be a integer number")
		os.Exit(-1)
	}
	if value <= 2 {
		fmt.Println("[-p / --population] must be greater than 2")
		os.Exit(-1)
	}
	optionP = int(value)
}

func main() {
	if len(os.Args)%2 != 1 {
		fmt.Println("invalid number of arguments:", len(os.Args))
		os.Exit(-1)
	}
	for i := 1; i < len(os.Args); i += 2 {
		switch os.Args[i] {
		case "-g":
			fallthrough
		case "--generations":
			parseGenerations(os.Args[i+1])
		case "-i":
			fallthrough
		case "--input":
			optionI = os.Args[i+1]
		case "-o":
			fallthrough
		case "--output":
			optionO = os.Args[i+1]
		case "-p":
			fallthrough
		case "--population":
			parsePopulation(os.Args[i+1])
		default:
			fmt.Println("invalid option:", os.Args[i])
			os.Exit(-1)
		}
	}
	inputs, outputs, err := readExamples(optionI)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	constants, instructions := evolve(optionG, optionP, inputs, outputs)
	err = writeProgram(optionO, constants, instructions)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
