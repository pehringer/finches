package main

import (
	"fmt"
	"os"
	"strconv"
)

func invalidArguments() {
	fmt.Print(`
  //>
 //)    f  i  n  c  h  e  s
/ ^

command format:
	finches [EXAMPLES_CSV_FILEPATH] [OPTION] . . .
command options:
	-d / --destination  [FILEPATH]
	-g / --generations  [NUMBER]
	-i / --individuals  [NUMBER]
hints:
	adjust the --generations and or --individuals counts
	if the resulting 'error' or number of 'instructions'
	is too high

	each line in the example.csv file must contains ONE to
	SIXTEEN example inputs followed by ONE expected output,
	for example abs.csv:
		-0.1,0.1
		2.3,2.3
		4.5,4.5
		-6.7,6.7
		-8.9,8.9
`)
}

func main() {
	if len(os.Args) < 2 {
		invalidArguments()
		return
	}
	source := os.Args[1]
	if len(os.Args)%2 != 0 {
		invalidArguments()
		return
	}
	if source == "-h" || source == "--help" {
		invalidArguments()
		return
	}
	destination := "function.go"
	generations := 2048
	individuals := 512
	for i := 2; i < len(os.Args); i += 2 {
		switch os.Args[i] {
		case "-d":
			fallthrough
		case "--destination":
			destination = os.Args[i+1]
		case "-g":
			fallthrough
		case "--generations":
			number, _ := strconv.ParseInt(os.Args[i+1], 10, 64)
			generations = int(number)
		case "-i":
			fallthrough
		case "--individuals":
			number, _ := strconv.ParseInt(os.Args[i+1], 10, 64)
			individuals = int(number)
		default:
			invalidArguments()
			return
		}
	}
	inputs, outputs, err := readExamples(source)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	constants, instructions := evolve(generations, individuals, inputs, outputs)
	fmt.Printf(" -> %s\n", destination)
	err = writeProgram(destination, len(inputs[0]), constants, instructions)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
