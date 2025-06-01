package io

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"github.com/pehringer/mapper/internal/types"
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

func ReadMappings(filename string) ([]types.Mapping, error) {
	values, err := readCSV(filename)
	if err != nil {
		return nil, err
	}
	if len(values[0]) < 2 {
		return nil, fmt.Errorf("expected 2 or more columns")
	}
	mappings := make([]types.Mapping, len(values)-1)
	for i := 1; i < len(values); i++ {
		columns, err := parseFloats(values[i])
		if err != nil {
			return nil, fmt.Errorf("row %d: %w", i+1, err)
		}
		output := len(columns) - 1
		mappings[i-1].Inputs = columns[:output]
		mappings[i-1].Output = columns[output]
	}
	return mappings, nil
}
