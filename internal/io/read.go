package io

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readCSV(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	values, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not read csv: %w", err)
	}
	return values, nil
}

func parseFloats(values []string) ([]float64, error) {
	result := make([]float64, len(values))
	for i := range values {
		trimmed := strings.TrimSpace(values[i])
		float, err := strconv.ParseFloat(trimmed, 64)
		if err != nil {
			return nil, fmt.Errorf("column %d: %w", i+1, err)
		}
		result[i] = float
	}
	return result, nil
}

func ReadMappings(filename string) ([][]float64, []float64, error) {
	values, err := readCSV(filename)
	if err != nil {
		return nil, nil, err
	}
	if len(values[0]) < 2 {
		return nil, nil, fmt.Errorf("expected 2 or more columns")
	}
	inputs := make([][]float64, len(values)-1)
	outputs := make([]float64, len(values)-1)
	for i := 1; i < len(values); i++ {
		columns, err := parseFloats(values[i])
		if err != nil {
			return nil, nil, fmt.Errorf("row %d: %w", i+1, err)
		}
		output := len(columns) - 1
		inputs[i-1] = columns[:output]
		outputs[i-1] = columns[output]
	}
	return inputs, outputs, nil
}
