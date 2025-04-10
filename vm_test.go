package fungen

import (
	"errors"
	"reflect"
	"testing"
)

func TestInputOutputNone(t *testing.T) {
	inputs := []float32{}
	outputs := SetInputs(inputs).GetOutputs()
	if !reflect.DeepEqual(inputs, outputs) {
		t.Errorf("incorrect outputs")
	}
}

func TestInputOutputSome(t *testing.T) {
	inputs := []float32{0, 1, 1, 2, 3, 5}
	outputs := SetInputs(inputs).GetOutputs()
	if !reflect.DeepEqual(inputs, outputs) {
		t.Errorf("incorrect outputs")
	}
}


func TestErrorsADD(t *testing.T) {
	inputs := []float32{1.2}
	m := SetInputs(inputs)
	err := m.Execute([5]byte{0x00, 0x00, 0x00, 0x00})
	if !errors.Is(err, ErrorStackUnderflow) {
		t.Errorf("incorrect error")
	}
	outputs := m.GetOutputs()
	if !reflect.DeepEqual(inputs, outputs) {
		t.Errorf("incorrect outputs")
	}
}
