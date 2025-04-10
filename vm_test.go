package fungen

import (
	"reflect"
	"testing"
)

func TestInputOutputNone(t *testing.T) {
	inputs := []float32{0, 1, 1, 2, 3, 5}
	outputs := SetInputs(inputs).GetOutputs()
	if !reflect.DeepEqual(inputs, outputs) {
		t.Errorf("inputs should match outputs")
	}
}
