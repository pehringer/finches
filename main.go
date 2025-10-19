package main

import (
	"fmt"
	"os"
	"strconv"
)

const help = `
  //>
 //)    f  i  n  c  h  e  s
/ ^

command format:
	finches [EXAMPLES_CSV_FILEPATH] [OPTION] . . .
command options:
	-d / --destination  [FILEPATH]
	-g / --generations  [NUMBER_GREATER_THAN_ZERO]
	-i / --individuals  [NUMBER_GREATER_THAN_THREE]
hints:
	adjust the --generations and or --individuals counts
	if the resulting 'error' or number of 'instructions'
	is too high

	each line in the example.csv file must contains ONE to
	EIGHT example inputs followed by ONE expected output,
	for example abs.csv:
		-0.1,0.1
		2.3,2.3
		4.5,4.5
		-6.7,6.7
		-8.9,8.9
`

func exitHelp(err error) {
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(help)
	os.Exit(1)
}

func exitError(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func parseMinimum(value string, minimum int64) (int, error) {
	result, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("'%s' must be an integer", value)
	}
	if result < minimum {
		return 0, fmt.Errorf("'%s' must be greater than or equal to %d", value, minimum)
	}
	return int(result), nil
}

func main() {
	if len(os.Args) < 2 {
		exitHelp(fmt.Errorf("missing examples csv filepath"))
	}
	source := os.Args[1]
	if len(os.Args)%2 != 0 {
		exitHelp(fmt.Errorf("'%s' missing argument", os.Args[len(os.Args)-1]))
	}
	if source == "-h" || source == "--help" {
		exitHelp(nil)
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
			number, err := parseMinimum(os.Args[i+1], 1)
			if err != nil {
				exitHelp(fmt.Errorf("-g / --generations %w", err))
			}
			generations = number
		case "-i":
			fallthrough
		case "--individuals":
			number, err := parseMinimum(os.Args[i+1], 3)
			if err != nil {
				exitHelp(fmt.Errorf("-i / --individuals %w", err))
			}
			individuals = number
		default:
			exitHelp(fmt.Errorf("'%s' unknown option", os.Args[i]))
		}
	}
	inputs, outputs, err := readExamples(source)
	if err != nil {
		exitError(err)
	}
	constants, instructions := evolve(generations, individuals, inputs, outputs)
	fmt.Printf(" -> %s\n", destination)
	err = writeProgram(destination, len(inputs[0]), constants, instructions)
	if err != nil {
		exitError(err)
	}
}
